package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pssilv/Blog-aggregator/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
  quantity := int32(2)

  if len(cmd.args) == 1 {
    number, err := strconv.Atoi(cmd.args[0])
    if err != nil {
      return fmt.Errorf("Failed to convert string to number: %w", err)
    }
    quantity = int32(number)
  }

  posts, err := s.db.GetPostsForUser(context.Background(), quantity)
  if err != nil {
    return fmt.Errorf("Failed to get the posts: %w", err)
  }

  fmt.Printf("User: %s with %d posts\n", user.Name, len(posts))
  printPosts(posts)

  return nil
}

func printPosts(posts []database.Post) {
  for _, post := range posts {
    fmt.Println("------------------")
    fmt.Printf("Post: %s\n", post.Title)
    fmt.Printf("Post description: %s\n", post.Description.String)
    fmt.Printf("Post publication date: %v\n", post.PublishedAt)
    fmt.Printf("Post last update %s\n", post.UpdatedAt)
    fmt.Printf("Record creation date: %v\n", post.CreatedAt)
    fmt.Println("------------------")
  }
}
