package parsing

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

func MewAquamarketParser() Parser {
	return &aquamarketParser{
		&helper{&http.Client{Timeout: time.Second * 10}},
		regexp.MustCompile(`^.*, (\d+ \* )?(\d+) (к?г)(\W.*)?$`),
	}
}

type aquamarketParser struct {
	helper *helper
	regexp *regexp.Regexp
}

func (p *aquamarketParser) ParseBuckwheats() ([]Buckwheat, error) {
	document, err := p.helper.readDocument("https://aquamarket.ua/uk/1099-grechka")
	if err != nil {
		return nil, err
	}
	buckwheats := make([]Buckwheat, 0)
	selection := document.Find("div.product:not(.product-not-available)")
	for i := range selection.Nodes {
		if buckwheat, err := p.parseBuckwheat(selection.Eq(i)); err != nil {
			log.WithField("url", buckwheat.URL).Error(err)
		} else {
			buckwheats = append(buckwheats, buckwheat)
		}
	}
	return buckwheats, nil
}

func (p *aquamarketParser) parseBuckwheat(selection *goquery.Selection) (Buckwheat, error) {
	var (
		buckwheat Buckwheat
		err       error
	)
	for _, node := range selection.Find("a.product-name[href][title]").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				buckwheat.URL = attribute.Val
			} else if attribute.Key == "title" {
				buckwheat.Title = attribute.Val
			}
		}
	}
	if buckwheat.URL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no url")
	}
	if buckwheat.Title == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no title")
	}
	for _, node := range selection.Find("div.thumb-list img[data-src]").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "data-src" {
				buckwheat.ImageURL = attribute.Val
			}
		}
	}
	if buckwheat.ImageURL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no image url")
	}
	var price string
	for _, node := range selection.Find("span.price.product-price[content]").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "content" {
				price = attribute.Val
			}
		}
	}
	buckwheat.Price, err = strconv.ParseFloat(price, 64)
	if err != nil {
		return buckwheat, fmt.Errorf("parsing: parser failed to parse price, %v", err)
	}
	if buckwheat.Price <= 0 {
		return buckwheat, fmt.Errorf("parsing: parser found non-positive price")
	}
	groups := p.regexp.FindStringSubmatch(buckwheat.Title)
	buckwheat.Mass, err = strconv.ParseFloat(groups[2], 64)
	if err != nil {
		return buckwheat, fmt.Errorf("parsing: parser failed to parse right mass, %v", err)
	}
	if buckwheat.Mass <= 0 {
		return buckwheat, fmt.Errorf("parsing: parser found non-positive right mass")
	}
	if groups[1] != "" {
		mass, err := strconv.ParseFloat(groups[1][:len(groups[1])-3], 64)
		if err != nil {
			return buckwheat, fmt.Errorf("parsing: parser failed to parse left mass, %v", err)
		}
		if mass <= 0 {
			return buckwheat, fmt.Errorf("parsing: parser found non-positive left mass")
		}
		buckwheat.Mass *= mass
	}
	switch groups[3] {
	case "кг":
		return buckwheat, nil
	case "г":
		buckwheat.Mass /= 1000
		return buckwheat, nil
	default:
		return buckwheat, fmt.Errorf("parsing: parser found unknown mass unit")
	}
}
