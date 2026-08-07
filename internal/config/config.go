package config

import (
	"encoding/json"
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

	content, err := os.ReadFile(homeDir + "/.gatorconfig.json")
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
