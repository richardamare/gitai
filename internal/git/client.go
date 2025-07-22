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

// GetUnifiedDiff returns the diff of all changes (staged and unstaged) with extended context
func (c *Client) GetUnifiedDiff() (string, error) {
    cmd := exec.Command("git", "diff", "-U50")
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to get unified diff. Is git installed? %w", err)
    }
    return strings.TrimSpace(string(output)), nil
}

// GetDiffFromMain returns the diff between the current branch and the main branch
func (c *Client) GetDiffFromMain(mainBranch string) (string, error) {
	cmd := exec.Command("git", "diff", mainBranch, "-U50")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff from %s branch. Is git installed and %s branch exists? %w", mainBranch, mainBranch, err)
	}
	return strings.TrimSpace(string(output)), nil
}

// GetBranchDiff returns the diff of the current branch compared to its upstream
func (c *Client) GetBranchDiff() (string, error) {
	cmd := exec.Command("git", "diff", "@{u}", "-U50")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get diff of current branch against upstream. Is git installed and is the branch tracked? %w", err)
	}
	return strings.TrimSpace(string(output)), nil
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