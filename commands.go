package main

import (
	"log"

	"github.com/pssilv/Blog-aggregator/internal/config"
)


type State struct {
  Cfg *config.Config
}

type Command struct {
  Name string
  Args []string
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

  handler, exists := c.CommandsName[cmd.Name]
  if !exists {
    log.Fatal("Command name doesn't exist")
  }

  if err := handler(s, cmd); err != nil {
    log.Fatal(err)
  }

  return nil
}

