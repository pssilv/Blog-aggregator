package main

import (
	"log"
	"os"
	"strings"

	"github.com/pssilv/Blog-aggregator/internal/config"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
    log.Fatal(err)
  }

  state := State {
    Cfg: &cfg,
  }

  commands := &Commands {
    CommandsName: make(map[string]func(*State, Command) error),
  }

  commands.Register("login", HandlerLogin)

  args := os.Args[1:]

  if len(args) < 2 {
    log.Fatal("Got less than 2 arguments")
  }


  command := Command {
    Name: strings.ToLower(args[0]),
    Args: args[1:],
  }

  if err := commands.Run(&state, command); err != nil {
    log.Fatal(err)
  }
}
