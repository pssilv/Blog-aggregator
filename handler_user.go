package main

import (
  "fmt"
  "errors"
  "log"
)

func HandlerLogin(s *State, cmd Command) error {
  if len(cmd.Args) == 0 {
    return errors.New("No arguments on command struct")
  }

  user := cmd.Args[0]

  if err := s.Cfg.SetUser(user); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("User: %v has been set\n", user)
  return nil
}

