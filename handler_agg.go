package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *State, cmd Command) error {
  RSSFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  fmt.Println(RSSFeed)

  return nil
}
