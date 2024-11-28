package main

import (
	"log"

	"github.com/pssilv/Blog-aggregator/internal/config"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

type State struct {
  cfg *config.Config
  db *database.Queries
  connection string
}

type Command struct {
  name string
  args []string
}

type Commands struct {
  CommandsName map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
  c.CommandsName[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
  if c.CommandsName == nil {
    log.Fatal("commandsName is  nil!")
  }

  handler, exists := c.CommandsName[cmd.name]
  if !exists {
    log.Fatal("Command name doesn't exist")
  }

  if err := handler(s, cmd); err != nil {
    log.Fatal(err)
  }

  return nil
}

