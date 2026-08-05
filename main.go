package main

import (
	"fmt"

	"github.com/davidheeren/gator/internal/config"
)

func main() {
	cfg, err := config.ReadFile()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	cfg.SetUser("david")

	cfg, err = config.ReadFile()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("db url: %s\n", cfg.DBUrl)
	fmt.Printf("current user name: %s\n", cfg.CurrentUserName)
}
