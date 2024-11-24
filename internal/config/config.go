package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
  home_directory, err := os.UserHomeDir()
  if err != nil {
    log.Fatal(err)
  }
  return fmt.Sprintf("%s/%s", home_directory, configFileName), nil
}

func write(cfg Config) error {
  filePath, err := getConfigFilePath()
  if err != nil {
    log.Fatal(err)
  }

  file, err := os.Open(filePath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  data, err := json.Marshal(cfg)
  if err != nil {
    log.Fatal(err)
  }


  if err := os.WriteFile(filePath, data, 0666); err != nil {
    log.Fatal(err)
  }

  return nil
}

func Read() Config {
  var config Config

  filePath, err := getConfigFilePath()
  if err != nil {
    log.Fatal(err)
  }

  file, err := os.Open(filePath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  data, err := os.ReadFile(filePath)

  if err := json.Unmarshal(data, &config); err != nil {
    log.Fatal(err)
  }

  return config
}

type Config struct {
  DbURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(user string) {
  c.CurrentUserName = user
  write(c)
}
