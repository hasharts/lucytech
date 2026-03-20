package scraper

import (
	"bytes"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	HTMLVersion string `json:"htmlVersion"`
	Title       string `json:"title"`
	Internal    int    `json:"internalLinks"`
	External    int    `json:"externalLinks"`
	HasLogin    bool   `json:"hasLogin"`
}

func ParseHTML(body io.Reader) (*Result, error) {

	// Read full HTML content
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	if err != nil {
		return nil, err
	}
	htmlStr := buf.String()

	// Detect HTML version from DOCTYPE
	htmlVersion := detectHTMLVersion(htmlStr)

	// Parse HTML using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if err != nil {
		return nil, err
	}

	res := &Result{}

	// Title
	res.Title = strings.TrimSpace(doc.Find("title").Text())

	// HTML Version
	res.HTMLVersion = htmlVersion

	// Links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}

		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
			res.External++
		} else {
			res.Internal++
		}
	})

	// Login form detection
	if doc.Find("input[type='password']").Length() > 0 {
		res.HasLogin = true
	}

	return res, nil
}

// Helper function to detect HTML version
func detectHTMLVersion(html string) string {
	lower := strings.ToLower(html)

	if strings.Contains(lower, "<!doctype html>") {
		return "HTML5"
	}

	if strings.Contains(lower, "html 4.01") {
		return "HTML 4.01"
	}

	if strings.Contains(lower, "xhtml") {
		return "XHTML"
	}

	return "Unknown"
}