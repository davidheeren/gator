package main

import (
	"fmt"
	"log"
	"os"

	"github.com/davidheeren/gator/internal/config"
)

type state struct {
	cfg *config.Config
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

	st := &state{
		cfg: cfg,
	}

	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

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
