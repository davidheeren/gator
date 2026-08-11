package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/davidheeren/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("unfollow command expects 'url' argument")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("no feed matches the url")
		}
		return err
	}

	feedFollow := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return err
	}

	fmt.Println("unfollowed", feed.Name)

	return nil
}
