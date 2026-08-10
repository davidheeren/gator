package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("users command expects 0 arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, u := range users {
		if u.Name == s.cfg.CurrentUserName {
			fmt.Println("*", u.Name, "(current)")
		} else {
			fmt.Println("*", u.Name)
		}
	}

	return nil
}
