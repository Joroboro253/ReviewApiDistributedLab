package main

import (
	"os"
	"review_api/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
