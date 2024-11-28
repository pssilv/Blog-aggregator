package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/pssilv/Blog-aggregator/internal/config"
	"github.com/pssilv/Blog-aggregator/internal/database"
)

func main() {
  cfg, err := config.Read()
  if err != nil {
    log.Fatal(err)
  }

  state := State {
    cfg: &cfg,
    connection: "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  }

  commands := &Commands {
    CommandsName: make(map[string]func(*State, Command) error),
  }

  db, err := sql.Open("postgres", state.connection)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  dbQueries := database.New(db)
  state.db = dbQueries


  commands.Register("login", HandlerLogin)
  commands.Register("register", HandlerRegister)
  commands.Register("reset", handlerReset)
  commands.Register("users", handlerList)

  args := os.Args[1:]

  if len(args) < 1 {
    log.Fatal("Missing arguments")
  }


  command := Command {
    name: strings.ToLower(args[0]),
    args: args[1:],
  }

  if err := commands.Run(&state, command); err != nil {
    log.Fatal(err)
  }

}
