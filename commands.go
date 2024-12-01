package main

import (
	"fmt"

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
    return fmt.Errorf("commandsName is  nil!")
  }

  handler, exists := c.CommandsName[cmd.name]
  if !exists {
    return fmt.Errorf("Command name doesn't exist")
  }

  if err := handler(s, cmd); err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  return nil
}

