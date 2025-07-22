package models

// CommitMessage represents a generated commit message
type CommitMessage struct {
    Message string `json:"message"`
}

// MrDetails represents MR information
type MrDetails struct {
    Title        string        `json:"title"`
    Description  string        `json:"description"`
    FileSummaries []FileSummary `json:"fileSummaries"`
}

// FileSummary represents a summary of changes in a file
type FileSummary struct {
    File        string `json:"file"`
    Description string `json:"description"`
}

// MrReviewDetails represents PR review feedback
type MrReviewDetails struct {
    Review []ReviewComment `json:"review"`
}

// ReviewComment represents a single review comment
type ReviewComment struct {
    File        string `json:"file"`
    Line        int    `json:"line"`
    Category    string `json:"category"`
    Comment     string `json:"comment"`
    CodeSnippet string `json:"codeSnippet"`
}

// MrTitle represents a PR title
type MrTitle struct {
    Title string `json:"title"`
}

// MrReviewSummary represents a general review summary for a PR
type MrReviewSummary struct {
    Summary string `json:"summary"`
}