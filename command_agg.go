package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/davidheeren/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("agg command expects 0 arguments")
	}

	feedUrl := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	fmt.Println("Title:", feed.Channel.Title)
	fmt.Println("Description:", feed.Channel.Description)
	fmt.Println("Link:", feed.Channel.Link)
	for _, item := range feed.Channel.Item {
		fmt.Println("    Title:", item.Title)
		desc := strings.ReplaceAll(item.Description, "\n", "\n      ")
		fmt.Println("    Description:")
		fmt.Println("      " + desc)
		fmt.Println("    PubDate:", item.PubDate)
		fmt.Println("    Link:", item.Link)
		fmt.Println()
	}

	return nil
}
