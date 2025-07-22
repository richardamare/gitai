package cmd

import (
	"fmt"

	"github.com/richardamare/gitai/internal/ai"
	"github.com/richardamare/gitai/internal/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCommitCommand creates the commit command
func NewCommitCommand() *cobra.Command {
    var autoCommit bool

    cmd := &cobra.Command{
        Use:   "commit",
        Short: "Generate AI-powered commit messages",
        Long:  "Generate commit messages using AI based on staged changes",
        RunE: func(cmd *cobra.Command, args []string) error {
            gitClient := git.NewClient()

            if !gitClient.IsGitRepo() {
                return fmt.Errorf("not in a git repository")
            }

            diff, err := gitClient.GetStagedDiff()
            if err != nil {
                return err
            }

            if diff == "" {
                return fmt.Errorf("no staged changes found")
            }

            apiKey := viper.GetString("openai_api_key")
            if apiKey == "" {
                return fmt.Errorf("OpenAI API key not found. Set OPENAI_API_KEY environment variable")
            }

            aiClient := ai.NewClient(apiKey)
            commitMsg, err := aiClient.GenerateCommitMessage(diff)
            if err != nil {
                return err
            }

            fmt.Printf("Generated commit message:\n%s\n\n", commitMsg.Message)

            if autoCommit {
                return gitClient.Commit(commitMsg.Message)
            }

            fmt.Println("Use --auto to automatically commit with this message")
            return nil
        },
    }

    cmd.Flags().BoolVarP(&autoCommit, "auto", "a", false, "Automatically commit with generated message")
    
    return cmd
}