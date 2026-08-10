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

	s.db.ResetUsers(context.Background())

	fmt.Println("users table has been reset")

	return nil
}
