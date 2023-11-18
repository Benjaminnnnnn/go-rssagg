package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	PubDate     string `xml:"pubDate"`
}

func urlToRssFeed(url string) (RSSFeed, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get(url)
	if err != nil {
		return RSSFeed{}, errors.New(fmt.Sprintf("Unable to get RSS feed from %v", url))
	}
	defer resp.Body.Close()

	// unmarshal response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, errors.New(fmt.Sprintf("Unable to read RSS feed response body %v", url))
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return RSSFeed{}, errors.New(fmt.Sprintf("Unable to unmarshal RSS feed response body %v", url))
	}

	return feed, nil
}
