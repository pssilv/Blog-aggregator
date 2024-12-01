package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
  if len(cmd.args) != 1 {
    return fmt.Errorf("Invalid number of arguments, expected: url")
  }
  url := cmd.args[0]

  feed, err := s.db.GetFeed(context.Background(), url)
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  feedFollowParams := database.CreateFeedFollowParams {
    ID: uuid.New(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
    UserID: user.ID,
    FeedID: feed.ID,
  }

  feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  printFeedFollow(feedFollow)
 
  return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowRow) {
  fmt.Printf("User: %v\n", feedFollow.UserName)
  fmt.Printf("%v feeds: %v\n", feedFollow.UserName, feedFollow.FeedsName)
}

func handlerListFollowingFeeds(s *State, cmd Command, user database.User) error {
  followingFeedUrls, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  for _, followingFeedUrl := range followingFeedUrls {
    feed, err := s.db.GetFeed(context.Background(), followingFeedUrl)

    if err != nil {
      return fmt.Errorf("Issue: %w", err)
    }

    fmt.Println(feed.Name)
  }

  return nil
}

func handlerUnfollow(s *State, cmd Command, user database.User) error {
  if len(cmd.args) != 1 {
    return fmt.Errorf("Invalid number of arguments, expected: url")
  }
  url := cmd.args[0]

  feed, err := s.db.GetFeed(context.Background(), url)
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  deleteFeedFollowParams := database.DeleteFeedFollowParams {
    UserID: user.ID,
    FeedID: feed.ID,
  }

  if err := s.db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams); err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  return nil
}
