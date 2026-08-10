package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("reset command expects 0 arguments")
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("users table has been reset")

	return nil
}
