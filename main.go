package main

import (
	"fmt"
	"os"
)

const version = "0.0.1"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fontgen <command>")
		fmt.Println("Commands:")
		fmt.Println("  version    Show version information")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println(version)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: version")
		os.Exit(1)
	}
}