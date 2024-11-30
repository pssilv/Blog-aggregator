package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

func fetchFeed(ctx context.Context, feedURL string) (RSSFeed, error) {
  var rssFeed RSSFeed

  client := http.Client{
    Timeout: 5 * time.Second,
  }

  req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
  if err != nil {
    return RSSFeed{}, fmt.Errorf("Error: %w", err)
  }

  req.Header.Add("gator", "User-Agent")

  res, err := client.Do(req)
  if err != nil {
    return RSSFeed{}, fmt.Errorf("Error: %w", err)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    error_message := fmt.Sprintf("Bad status code: %d\n", res.StatusCode)

    return RSSFeed{}, errors.New(error_message)
  }

  body, err := io.ReadAll(res.Body)

  if err := xml.Unmarshal(body, &rssFeed); err != nil {
    return RSSFeed{}, fmt.Errorf("Error: %w", err)
  }

  rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
  rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
  for idx, item := range rssFeed.Channel.Item {
    item.Title = html.UnescapeString(item.Title)
    item.Description = html.UnescapeString(item.Description)

    rssFeed.Channel.Item[idx] = item
  }

    
  return rssFeed, nil
}

type RSSFeed struct {
  Channel struct {
    Title       string `xml:"title"`
    Link        string `xml:"link"`
    Description string `xml:"description"`
    Item        []RSSItem `xml:"item"`
  } `xml:"channel"`
}

type RSSItem struct {
  Title       string `xml:"title"`
  Link        string `xml:"link"`
  Description string `xml:"description"`
  PubDate     string `xml:"pubDate"`
}
