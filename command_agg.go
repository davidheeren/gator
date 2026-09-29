package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/davidheeren/gator/internal/database"
	"github.com/davidheeren/gator/internal/rss"
	"github.com/google/uuid"
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

		_, err = s.db.MarkFeedFetched(context.Background(), feed.ID)
		if err != nil {
			return err
		}

		feedData, err := rss.FetchFeed(context.Background(), feed.Url)
		if err != nil {
			return err
		}

		for _, item := range feedData.Channel.Item {
			nullDescription := sql.NullString{
				String: item.Description,
				Valid:  item.Description != "",
			}

			// need to handle different date layouts
			dateLayouts := []string{time.RFC1123Z, time.RFC1123}
			var nullPubDate sql.NullTime
			for _, layout := range dateLayouts {
				pubDate, err := time.Parse(layout, item.PubDate)
				nullPubDate = sql.NullTime{
					Time:  pubDate,
					Valid: err == nil,
				}
				if nullPubDate.Valid {
					break
				}
			}

			postArgs := database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       item.Title,
				Url:         item.Link,
				Description: nullDescription,
				PublishedAt: nullPubDate,
				FeedID:      feed.ID,
			}
			_, err := s.db.CreatePost(context.Background(), postArgs)
			if err != nil {
				if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
					// fmt.Printf("post %s already exits\n", item.Title)
				} else {
					fmt.Printf("Couldn't create post %s: %v", item.Title, err)
				}
			}
		}
	}

	return nil
}
