package main

import (
	_ "github.com/lib/pq"

	"database/sql"
	"log"
	"os"
	"strings"
  "fmt"
  "context"

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
    connection: cfg.DBURL,
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
  commands.Register("agg", handlerAggregate)
  commands.Register("feeds", handlerListFeeds)

  commands.Register("addfeed", middlewareLoggedIn(handlerFeed))
  commands.Register("follow", middlewareLoggedIn(handlerFollow))
  commands.Register("unfollow", middlewareLoggedIn(handlerUnfollow))
  commands.Register("following", middlewareLoggedIn(handlerListFollowingFeeds))
  commands.Register("browse", middlewareLoggedIn(handlerBrowse))

  args := os.Args[1:]

  if len(args) < 1 {
    log.Fatal("Missing command")
  }


  command := Command {
    name: strings.ToLower(args[0]),
    args: args[1:],
  }

  if err := commands.Run(&state, command); err != nil {
    log.Fatal(err)
  }

}

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {  
  return func(s *State, cmd Command) error {
    currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
      return fmt.Errorf("Issue: %w", err)
    }

    return handler(s, cmd, currentUser)
  }
}
