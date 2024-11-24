package main

import (
	"fmt"

	"github.com/pssilv/Blog-aggregator/internal/config"
)

func main() {
  cfg := config.Read()

  cfg.SetUser("pssilv")

  fmt.Println(cfg)
}
