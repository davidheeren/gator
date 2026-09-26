package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/davidheeren/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("agg command expects 'request_duration' (ie: 1s, 1m, 1h) argument")
	}

	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", time_between_reqs.String())
	ticker := time.NewTicker(time_between_reqs)

	for _ = range ticker.C {
		feed, err := s.db.GetNextFeedToFetch(context.Background())
		if err != nil {
			return err
		}

		fmt.Printf("\n----%s----\n\n", feed.Name)
		_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
		if err != nil {
			return err
		}

		feedData, err := rss.FetchFeed(context.Background(), feed.Url)
		if err != nil {
			return err
		}

		for _, item := range feedData.Channel.Item {
			fmt.Println(item.Title)
		}
	}

	return nil
}
