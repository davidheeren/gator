package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("feeds command expects 0 arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println("Name:", feed.Name)
		fmt.Println("Url:", feed.Url)
		fmt.Println("User name:", feed.UserName)
		fmt.Println()
	}

	return nil
}
