package cmd

import (
	"errors"
	"fmt"
	"os"
	"unicode"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/zquestz/go-chatgpt-twitter-bot/bot"
)

const (
	appName               = "go-chatgpt-twitter-bot"
	version               = "0.0.1"
	twitterBearerTokenEnv = "TWITTER_BEARER_TOKEN"
)

// Stores configuration data.
var config Config

// SearchCmd is the main command for Cobra.
var RunCmd = &cobra.Command{
	Use:   "go-chatgpt-twitter-bot <handle>",
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

	err = godotenv.Load()
	if err != nil {
		bail(fmt.Errorf("Error loading .env file: %s", err))
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
	RunCmd.PersistentFlags().StringVarP(
		&config.Handle, "handle", "", "", "twitter handle")
	RunCmd.PersistentFlags().StringVarP(
		&config.UserID, "userid", "", "", "twitter userid")
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

	if config.Handle == "" {
		return errors.New("handle is required")
	}

	if config.UserID == "" {
		return errors.New("userid is required")
	}

	if len(args) != 0 {
		// Don't return an error, help screen is more appropriate.
		help := cmd.HelpFunc()
		help(cmd, args)
		return nil
	}

	bearerToken := os.Getenv(twitterBearerTokenEnv)
	if bearerToken == "" {
		return errors.New("TWITTER_BEARER_TOKEN env var is required")
	}

	err := bot.Run(config.UserID, bearerToken)
	if err != nil {
		return err
	}

	return nil
}
