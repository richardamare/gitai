package models

// CommitMessage represents a generated commit message
type CommitMessage struct {
    Message string `json:"message"`
}

// PrDetails represents PR information
type PrDetails struct {
    Title        string        `json:"title"`
    Description  string        `json:"description"`
    FileSummaries []FileSummary `json:"fileSummaries"`
}

// FileSummary represents a summary of changes in a file
type FileSummary struct {
    File        string `json:"file"`
    Description string `json:"description"`
}

// PrReviewDetails represents PR review feedback
type PrReviewDetails struct {
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

// PrTitle represents a PR title
type PrTitle struct {
    Title string `json:"title"`
}