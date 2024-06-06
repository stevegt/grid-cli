package main

import (
	"fmt"
	"os"

	. "github.com/stevegt/goadapt"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: grid {subcommand} [args...]")
		fmt.Println("       grid --show {subcommand}")
		os.Exit(1)
	}

	ensureDirectories()
	loadPeers()
	connectToPeers()

	switch args[1] {
	case "--show":
		if len(args) < 3 {
			fmt.Println("Usage: grid --show {subcommand}")
			os.Exit(1)
		}
		subcommand := args[2]
		showPromise(subcommand)
	case "start-server":
		startWebSocketServer()
	default:
		subcommand := args[1]
		executeSubcommand(subcommand, args[2:])
	}
}
