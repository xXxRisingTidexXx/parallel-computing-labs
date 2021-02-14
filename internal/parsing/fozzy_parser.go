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

func NewFozzyParser() Parser {
	return &fozzyParser{
		&helper{&http.Client{Timeout: time.Second * 15}},
		regexp.MustCompile(`Фасовка: (\d+\*)?(\d+)(к?г)`),
	}
}

type fozzyParser struct {
	helper *helper
	regexp *regexp.Regexp
}

func (p *fozzyParser) ParseBuckwheats() ([]Buckwheat, error) {
	document, err := p.helper.readDocument("https://fozzyshop.ua/300143-krupa-grechnevaya")
	if err != nil {
		return nil, err
	}
	buckwheats := make([]Buckwheat, 0)
	selection := document.Find(
		"div#js-product-list div.js-product-miniature-wrapper:not(.product_grey)",
	)
	for i := range selection.Nodes {
		if buckwheat, err := p.parseBuckwheat(selection.Eq(i)); err != nil {
			log.WithField("url", buckwheat.URL).Error(err)
		} else {
			buckwheats = append(buckwheats, buckwheat)
		}
	}
	return buckwheats, nil
}

func (p *fozzyParser) parseBuckwheat(selection *goquery.Selection) (Buckwheat, error) {
	var (
		buckwheat Buckwheat
		err       error
	)
	for _, node := range selection.Find("a.thumbnail").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				buckwheat.URL = attribute.Val
			}
		}
	}
	if buckwheat.URL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no url")
	}
	for _, node := range selection.Find("img.product-thumbnail-first").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "src" {
				buckwheat.ImageURL = attribute.Val
			}
		}
	}
	if buckwheat.ImageURL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no image url")
	}
	buckwheat.Title = selection.Find("div.product-title > a").Text()
	if buckwheat.Title == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no title")
	}
	var price string
	for _, node := range selection.Find("span.product-price").Nodes {
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
	groups := p.regexp.FindStringSubmatch(
		selection.Find("div.product-reference > a").First().Text(),
	)
	buckwheat.Mass, err = strconv.ParseFloat(groups[2], 64)
	if err != nil {
		return buckwheat, fmt.Errorf("parsing: parser failed to parse right mass, %v", err)
	}
	if buckwheat.Mass <= 0 {
		return buckwheat, fmt.Errorf("parsing: parser found non-positive right mass")
	}
	if groups[1] != "" {
		mass, err := strconv.ParseFloat(groups[1][:len(groups[1])-1], 64)
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
