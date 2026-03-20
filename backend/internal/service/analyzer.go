package service

import (
    "net/http"

    "web-analyzer/internal/scraper"
)

func Analyze(url string) (*scraper.Result, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    result, err := scraper.ParseHTML(resp.Body)
    if err != nil {
        return nil, err
    }

    return result, nil
}