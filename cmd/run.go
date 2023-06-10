package cmd

import (
	"fmt"
	"os"
	"unicode"

	"github.com/spf13/cobra"
)

const (
	appName = "go-chatgpt-twitter-bot"
	version = "0.0.1"
)

// Stores configuration data.
var config Config

// SearchCmd is the main command for Cobra.
var RunCmd = &cobra.Command{
	Use:   "go-chatgpt-twitter-bot",
	Short: "Twitter bot backed by ChatGPT",
	Long:  `Twitter bot backed by ChatGPT.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := performCommand(cmd, args)
		if err != nil {
			bail(err)
		}
	},
}

func init() {
	err := config.Load()
	if err != nil {
		bail(fmt.Errorf("failed to load configuration: %s", err))
	}

	prepareFlags()
}

func bail(err error) {
	fmt.Fprintf(os.Stderr, "[Error] %s\n", capitalize(err.Error()))
	os.Exit(1)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return ""
	}
	tmp := []rune(str)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func completion(cmd *cobra.Command, c string) {
	switch c {
	case "bash":
		err := cmd.GenBashCompletion(os.Stdout)
		if err != nil {
			bail(fmt.Errorf("failed to generate bash completion: %w", err))
		}
	case "zsh":
		if err := cmd.GenZshCompletion(os.Stdout); err != nil {
			bail(fmt.Errorf("failed to generate zsh completion: %w", err))
		}
	case "fish":
		if err := cmd.GenFishCompletion(os.Stdout, true); err != nil {
			bail(fmt.Errorf("failed to generate fish completion: %w", err))
		}
	case "powershell":
		err := cmd.GenPowerShellCompletion(os.Stdout)
		if err != nil {
			bail(fmt.Errorf("failed to generate powershell completion: %w", err))
		}
	default:
		bail(fmt.Errorf("completion not supported: %s", c))
	}
}

func prepareFlags() {
	RunCmd.PersistentFlags().BoolVarP(
		&config.DisplayVersion, "version", "", false, "display version")
	RunCmd.PersistentFlags().BoolVarP(
		&config.Verbose, "verbose", "v", config.Verbose, "verbose mode")
	RunCmd.PersistentFlags().StringVarP(
		&config.Completion, "completion", "", "", "completion script for bash, zsh, fish or powershell")
}

// Where all the work happens.
func performCommand(cmd *cobra.Command, args []string) error {
	if config.DisplayVersion {
		fmt.Printf("%s %s\n", appName, version)
		return nil
	}

	if config.Completion != "" {
		completion(cmd, config.Completion)
		return nil
	}

	return nil
}
