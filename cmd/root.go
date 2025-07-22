package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
    Use:   "gitai",
    Short: "AI-powered Git CLI tool",
    Long:  "A CLI tool that uses AI to generate commit messages, PR descriptions, and code reviews",
}

// Execute runs the root command
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)
    
    // Add subcommands
    rootCmd.AddCommand(NewCommitCommand())
    // Add other commands here: PR, review, etc.
}

func initConfig() {
    viper.AutomaticEnv()
    viper.SetEnvPrefix("GITAI")
    
    // Bind environment variables
    viper.BindEnv("openai_api_key", "OPENAI_API_KEY")
}