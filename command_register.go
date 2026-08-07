package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davidheeren/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("register command expects 'name' argument")
	}

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	createdUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	s.cfg.SetUser(createdUser.Name)

	fmt.Println("new user has been registered and logged in")
	fmt.Println("id:", createdUser.ID)
	fmt.Println("CreatedAt:", createdUser.CreatedAt)
	fmt.Println("UpdatedAt:", createdUser.UpdatedAt)
	fmt.Println("Name:", createdUser.Name)
	return nil
}
