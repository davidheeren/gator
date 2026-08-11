package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davidheeren/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return errors.New("addfeed command expects 'name, url' arguments")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	createdFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return err
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    createdFeed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return err
	}

	fmt.Println("new feed has been created and followed by", s.cfg.CurrentUserName)
	fmt.Println("Id:", createdFeed.ID)
	fmt.Println("CreatedAt:", createdFeed.CreatedAt)
	fmt.Println("UpdatedAt:", createdFeed.UpdatedAt)
	fmt.Println("Name:", createdFeed.Name)
	fmt.Println("Url:", createdFeed.Url)
	fmt.Println("UserID:", createdFeed.UserID)

	return nil
}
