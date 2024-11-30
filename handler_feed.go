package main

import (
  "errors"
  "context"
  "fmt"
  "time"

  "github.com/google/uuid"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

func handlerFeed(s *State, cmd Command) error {
  if len(cmd.args) < 2 {
    return errors.New("Insufficient arguments")
  }
  name := cmd.args[0]
  url := cmd.args[1]


  feedExists, _ := s.db.GetFeed(context.Background(), url)
  if feedExists.Url != "" {
    return fmt.Errorf("Feed url: %s already exists", feedExists.Url)
  }

  currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
  if err != nil {
    return fmt.Errorf("Error: %w", err)
  }
  
  feedParams := database.CreateFeedParams{
    ID: uuid.New(), 
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
    Name: name,
    Url: url,
    UserID: currentUser.ID,
  }
 
  feed, err := s.db.CreateFeed(context.Background(), feedParams)
  if err != nil {
    return fmt.Errorf("Error: %w", err)
  }

  fmt.Printf("Feed: %v has been added\n", feed.Name)
  printFeed(feed)

  return nil
}

func printFeed(feed database.Feed) {
  fmt.Println("-------- Feed content: --------")
  fmt.Printf("ID: %v\n", feed.ID)
  fmt.Printf("CreatedAt: %v\n", feed.CreatedAt)
  fmt.Printf("UpdatedAt: %v\n", feed.Name)
  fmt.Printf("Name: %v\n", feed.Url)
  fmt.Printf("UserID: %v\n", feed.UserID)
  fmt.Println("-------------------------------")
}

func handlerListFeeds(s *State, cmd Command) error {
  rows, err := s.db.ListFeeds(context.Background())
  if err != nil {
    return fmt.Errorf("Error: %w", err)
  }

  for _, row := range rows {
    fmt.Println("-------------------------------")
    fmt.Printf("Feed name: %v\n", row.FeedName)
    fmt.Printf("Feed url: %v\n", row.FeedUrl)
    fmt.Printf("Creator name: %v\n", row.UserName.String)
    fmt.Println("-------------------------------")
  } 

  return nil
}
