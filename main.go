package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/davidheeren/gator/internal/config"
	"github.com/davidheeren/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db *database.Queries
}

type command struct {
	name string
	args []string
}

func main() {
	cfg, err := config.ReadFile()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	dbQueries := database.New(db)

	st := &state{
		cfg: cfg,
		db: dbQueries,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)

	if len(os.Args) < 2 {
		printHelp(cmds)
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(st, cmd)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func printHelp(cmds commands) {
	fmt.Println("Valid gator commands:")
	for n := range cmds.handlers {
		fmt.Println(n)
	}
}
