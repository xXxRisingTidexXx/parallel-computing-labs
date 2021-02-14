package parsing

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type helper struct {
	client *http.Client
}

func (h *helper) readDocument(url string) (*goquery.Document, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("parsing: helper failed to make a request, %v", err)
	}
	response, err := h.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("parsing: helper failed to send a request, %v", err)
	}
	if response.StatusCode != http.StatusOK {
		_ = response.Body.Close()
		return nil, fmt.Errorf("parsing: helper failed with status %s", response.Status)
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		_ = response.Body.Close()
		return nil, fmt.Errorf("parsing: helper failed to read a document, %v", err)
	}
	if err := response.Body.Close(); err != nil {
		return nil, fmt.Errorf("parsing: helper failed to close a response body, %v", err)
	}
	return document, nil
}
