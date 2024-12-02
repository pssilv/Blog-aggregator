package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

func handlerAggregate(s *State, cmd Command) error {
  if len(cmd.args) != 1 {
    return fmt.Errorf("Invalid number of arguments, expected: time_between_reqs")
  }
  timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  ticker := time.NewTicker(timeBetweenReqs)

  fmt.Printf("Collecting feeds every: %v\n", timeBetweenReqs)

  for ; ; <-ticker.C {
    scrapeFeeds(s)
  }
}

func scrapeFeeds(s *State) {
  feed, err := s.db.GetNextFeedToFetch(context.Background())
  if err != nil {
    log.Printf("Couldn't get the next feed to fetch: %v\n", err)
  }

  log.Println("Found a feed to fetch")
  scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
  _, err := db.MarkFeedFetched(context.Background(), feed.ID)
  if err != nil {
    log.Printf("Couldn't mark feed: %s, fetched: %v\n", feed.Name, err)
  }

  feedData, err := fetchFeed(context.Background(), feed.Url)
  if err != nil {
    log.Printf("Couldn't collect feed: %s: %v\n", feed.Name, err)
  }

  for _, item := range feedData.Channel.Item {
    fmt.Println("--------------------")
    fmt.Printf("Item: %s\n", item.Title)
    fmt.Printf("Description: %s\n", item.Description)
    fmt.Printf("Publication date: %s\n", item.PubDate)
    fmt.Printf("Link: %s\n", item.Link)
    fmt.Println("--------------------")
  }

  postParams := database.CreatePostParams{
    ID: uuid.New(),
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    Title: feedData.Channel.Title,
    Url: feed.Url,
    Description: sql.NullString{String: feedData.Channel.Description, Valid: true},
    PublishedAt: feed.CreatedAt,
    FeedID: feed.ID,
  }

  _, err = db.CreatePost(context.Background(), postParams)
  if err, ok := err.(*pq.Error); ok {
    if err.Code != "23505" {
      log.Printf("Failed to create a post: %v\n", err)
    }
  }
  log.Printf("Collected feed: %v with %v posts found\n", feedData.Channel.Title, len(feedData.Channel.Item))
}
