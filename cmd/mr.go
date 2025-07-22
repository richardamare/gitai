package cmd

import (
	"fmt"
	"os"

	"github.com/richardamare/gitai/internal/ai"
	"github.com/richardamare/gitai/internal/git"
	"github.com/spf13/cobra"
)

func NewMRCommand() *cobra.Command {
	mrCmd := &cobra.Command{
		Use:   "mr",
		Short: "AI-powered Merge Request tools",
		Long:  "A CLI tool that uses AI to generate merge request descriptions and titles.",
	}

	mrCmd.AddCommand(NewMRTitleCommand())
	mrCmd.AddCommand(NewMRReviewCommand())
	mrCmd.AddCommand(NewMRDetailsCommand())

	return mrCmd
}

func NewMRReviewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "review",
		Short: "Generate a review for the current merge request",
		Long:  "This command generates a review for the current merge request based on the git diff of the current branch.",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitClient := git.NewClient()
			diff, err := gitClient.GetDiffFromMain("master")
			if err != nil {
				return fmt.Errorf("failed to get git diff of current branch: %w", err)
			}

			if diff == "" {
				fmt.Println("No changes found on current branch.")
				return nil
			}

			aiClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))
			reviewDetails, err := aiClient.ReviewMR(diff)
			if err != nil {
				return fmt.Errorf("failed to generate MR review from AI: %w", err)
			}

			fmt.Println("AI Review:")
			for _, review := range reviewDetails.Review {
				fmt.Printf("\nFile: %s:%d\n", review.File, review.Line)
				fmt.Printf("Category: %s\n", review.Category)
				fmt.Printf("Comment: %s\n", review.Comment)
				if review.CodeSnippet != "" {
					fmt.Printf("Code Snippet:\n```\n%s\n```\n", review.CodeSnippet)
				}
			}
			return nil
		},
	}
}

func NewMRTitleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "title",
		Short: "Generate a title for the current merge request",
		Long:  "This command generates a title for the current merge request based on the git diff from the main branch.",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitClient := git.NewClient()
			diff, err := gitClient.GetDiffFromMain("master")
			if err != nil {
				return fmt.Errorf("failed to get git diff from master branch: %w", err)
			}

			if diff == "" {
				fmt.Println("No changes found compared to main branch.")
				return nil
			}

			aiClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))
			title, err := aiClient.GenerateMRTitle(diff)
			if err != nil {
				return fmt.Errorf("failed to generate MR title from AI: %w", err)
			}

			fmt.Printf("Generated MR Title: %s\n", title)
			return nil
		},
	}
}

func NewMRDetailsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "details",
		Short: "Generate a description for the current merge request",
		Long:  "This command generates a description for the current merge request based on the git diff from the main branch.",
		RunE: func(cmd *cobra.Command, args []string) error {
			gitClient := git.NewClient()
			diff, err := gitClient.GetDiffFromMain("master")
			if err != nil {
				return fmt.Errorf("failed to get git diff from master branch: %w", err)
			}

			if diff == "" {
				fmt.Println("No changes found compared to main branch.")
				return nil
			}

			aiClient := ai.NewClient(os.Getenv("OPENAI_API_KEY"))
			details, err := aiClient.GenerateMRDetails(diff)
			if err != nil {
				return fmt.Errorf("failed to generate MR details from AI: %w", err)
			}

			fmt.Printf("Generated MR Title: %s\n", details.Title)
			fmt.Printf("Generated MR Description: %s\n", details.Description)
			fmt.Println("--------------------------------")
			for _, file := range details.FileSummaries {
				fmt.Printf("File: %s\n", file.File)
				fmt.Printf("Description: %s\n", file.Description)
				fmt.Println("--------------------------------")
			}
			return nil
		},
	}
}