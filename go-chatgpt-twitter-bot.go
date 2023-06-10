package main

import (
	"fmt"
	"os"

	"github.com/zquestz/go-chatgpt-twitter-bot/cmd"
)

func main() {
	setupSignalHandlers()

	if err := cmd.RunCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
