package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// Client handles git operations
type Client struct{}

// NewClient creates a new git client
func NewClient() *Client {
    return &Client{}
}

// GetStagedDiff returns the staged diff with extended context
func (c *Client) GetStagedDiff() (string, error) {
    cmd := exec.Command("git", "diff", "--staged", "-U50")
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to get staged diff. Is git installed? %w", err)
    }
    return strings.TrimSpace(string(output)), nil
}

// GetDiff returns the diff for specified files or all changes
func (c *Client) GetDiff(files ...string) (string, error) {
    args := []string{"diff"}
    if len(files) > 0 {
        args = append(args, files...)
    }
    
    cmd := exec.Command("git", args...)
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to get diff: %w", err)
    }
    return strings.TrimSpace(string(output)), nil
}

// Commit creates a commit with the given message
func (c *Client) Commit(message string) error {
    cmd := exec.Command("git", "commit", "-m", message)
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("failed to commit: %w", err)
    }
    fmt.Println("✅ Successfully committed changes!")
    return nil
}

// GetCurrentBranch returns the current git branch
func (c *Client) GetCurrentBranch() (string, error) {
    cmd := exec.Command("git", "branch", "--show-current")
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to get current branch: %w", err)
    }
    return strings.TrimSpace(string(output)), nil
}

// IsGitRepo checks if current directory is a git repository
func (c *Client) IsGitRepo() bool {
    cmd := exec.Command("git", "rev-parse", "--git-dir")
    return cmd.Run() == nil
}