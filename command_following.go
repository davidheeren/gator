package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("following command expects no arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
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
