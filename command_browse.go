package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/davidheeren/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.args) == 1 {
		limit, err = strconv.Atoi(cmd.args[0])
		if err != nil {
			return errors.New("optional 'limit' (default 2) argument must be an integer")
		}
	} else if len(cmd.args) != 0 {
		return errors.New("browse command takes optional 'limit' (default 2) argument")
	}

	getPostsArgs := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsArgs)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
