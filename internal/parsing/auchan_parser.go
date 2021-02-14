package parsing

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NewAuchanParser() Parser {
	return &auchanParser{&helper{&http.Client{Timeout: time.Second * 10}}}
}

type auchanParser struct {
	helper *helper
}

func (p *auchanParser) ParseBuckwheats() ([]Buckwheat, error) {
	document, err := p.helper.readDocument(
		"https://auchan.zakaz.ua/uk/categories/buckwheat-auchan/",
	)
	if err != nil {
		return nil, err
	}
	buckwheats, selection := make([]Buckwheat, 0), document.Find("a.product-tile")
	for i := range selection.Nodes {
		if buckwheat, err := p.parseBuckwheat(selection.Eq(i)); err != nil {
			log.WithField("url", buckwheat.URL).Error(err)
		} else {
			buckwheats = append(buckwheats, buckwheat)
		}
	}
	return buckwheats, nil
}

func (p *auchanParser) parseBuckwheat(selection *goquery.Selection) (Buckwheat, error) {
	var (
		buckwheat Buckwheat
		err       error
	)
	for _, attribute := range selection.Nodes[0].Attr {
		if attribute.Key == "href" {
			buckwheat.URL = attribute.Val
		} else if attribute.Key == "title" {
			buckwheat.Title = attribute.Val
		}
	}
	if buckwheat.URL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no url")
	}
	buckwheat.URL = "https://auchan.zakaz.ua" + buckwheat.URL
	if buckwheat.Title == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no title")
	}
	for _, node := range selection.Find("img.product-tile__image-i").Nodes {
		for _, attribute := range node.Attr {
			if attribute.Key == "src" {
				buckwheat.ImageURL = attribute.Val
			}
		}
	}
	if buckwheat.ImageURL == "" {
		return buckwheat, fmt.Errorf("parsing: parser found no image url")
	}
	buckwheat.Price, err = strconv.ParseFloat(
		selection.Find("span.Price__value_caption").Text(),
		64,
	)
	if err != nil {
		return buckwheat, fmt.Errorf("parsing: parser failed to parse price, %v", err)
	}
	if buckwheat.Price <= 0 {
		return buckwheat, fmt.Errorf("parsing: parser found non-positive price")
	}
	words := strings.Split(selection.Find("div.product-tile__weight").Text(), " ")
	if len(words) < 2 || len(words) > 3 {
		return buckwheat, fmt.Errorf("parsing: parser failed to split price literal")
	}
	buckwheat.Mass, err = strconv.ParseFloat(words[len(words)-2], 64)
	if err != nil {
		return buckwheat, fmt.Errorf("parsing: parser failed to parse mass, %v", err)
	}
	if buckwheat.Mass <= 0 {
		return buckwheat, fmt.Errorf("parsing: parser found non-positive mass")
	}
	switch words[len(words)-1] {
	case "кг":
		return buckwheat, nil
	case "г":
		buckwheat.Mass /= 1000
		return buckwheat, nil
	default:
		return buckwheat, fmt.Errorf("parsing: parser found unknown mass unit")
	}
}
