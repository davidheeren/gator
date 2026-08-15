package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DBUrl           string `json:"db_json"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadFile() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	path := homeDir + "/.gatorconfig.json"

	if _, err := os.Stat(path); err != nil {
		return nil, errors.New("No ~.gatorconfig.json. Please enter the db_json into the config file")
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	content, err := json.Marshal(c)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.WriteFile(homeDir+"/.gatorconfig.json", content, 0644)
	if err != nil {
		return err
	}

	return nil
}
