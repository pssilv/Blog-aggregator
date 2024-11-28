package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

func HandlerLogin(s *State, cmd Command) error {
  if len(cmd.args) == 0 {
    return errors.New("Missing user")
  }

  name := cmd.args[0]

  userExists, _ := s.db.GetUser(context.Background(), name)
  if userExists == (database.User{}) {
    log.Fatal("User doesn't exist")
  }

  if err := s.cfg.SetUser(name); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("User: %v has been set\n", name)
  return nil
}

func HandlerRegister(s *State, cmd Command) error {
  if len(cmd.args) == 0 {
    return errors.New("Missing user")
  }

  name := cmd.args[0]

  userExists, _ := s.db.GetUser(context.Background(), name)
  if userExists.Name != "" {
    log.Fatalf("User: %s already exist", userExists.Name)
  }

  userParams := database.CreateUserParams {
    ID: uuid.New(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
    Name: name,
  }

  user, err := s.db.CreateUser(context.Background(), userParams)
  if err != nil {
    fmt.Println("UUID:", userParams.ID)
    log.Fatal(err)
  }

  if err := s.cfg.SetUser(user.Name); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("User: %v has been created\n", user.Name)
  fmt.Println(user)

  return nil
}

func handlerReset(s *State, cmd Command) error {
  if err := s.db.ResetUsers(context.Background()); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Table: Users have been reset\n")
 
  return nil
}


func handlerList(s * State, cmd Command) error {
  users, err := s.db.ListUsers(context.Background())
  if err != nil {
    log.Fatal(err)
  }

  for _, user := range users {
    if user == s.cfg.CurrentUserName {
      fmt.Printf("* %s (current)\n", user)
    } else {
    fmt.Printf("* %s\n", user)
    }
  }

  return nil
}
