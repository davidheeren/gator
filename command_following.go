package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/davidheeren/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return errors.New("following command expects no arguments")
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println(s.cfg.CurrentUserName, "is following:")
	for _, f := range follows {
		fmt.Println(f.FeedName)
	}

	return nil
}
