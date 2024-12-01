package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
  home_directory, err := os.UserHomeDir()
  if err != nil {
     return "", fmt.Errorf("Issue: %w", err)
  }
  return fmt.Sprintf("%s/%s", home_directory, configFileName), nil
}

func write(cfg Config) error {
  filePath, err := getConfigFilePath()
  if err != nil {
    return fmt.Errorf("Issue: %w", err)
  }

  file, err := os.Open(filePath)
  if err != nil {
     return fmt.Errorf("Issue: %w", err)
  }
  defer file.Close()

  data, err := json.Marshal(cfg)
  if err != nil {
     return fmt.Errorf("Issue: %w", err)
  }


  if err := os.WriteFile(filePath, data, 0666); err != nil {
     return fmt.Errorf("Issue: %w", err)
  }

  return nil
}

func Read() (Config, error) {
  var config Config

  filePath, err := getConfigFilePath()
  if err != nil {
     return Config{}, fmt.Errorf("Issue: %w", err)
  }

  file, err := os.Open(filePath)
  if err != nil {
     return Config{}, fmt.Errorf("Issue: %w", err)
  }
  defer file.Close()

  data, err := os.ReadFile(filePath)

  if err := json.Unmarshal(data, &config); err != nil {
     return Config{}, fmt.Errorf("Issue: %w", err)
  }

  return config, nil
}

type Config struct {
  DbURL           string `json:"db_url"`
  CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
  c.CurrentUserName = userName
  return write(*c)
}
