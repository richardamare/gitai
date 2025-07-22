package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/richardamare/gitai/internal/models"
	"github.com/sashabaranov/go-openai"
)

// Client handles AI operations
type Client struct {
    client *openai.Client
}

// NewClient creates a new AI client
func NewClient(apiKey string) *Client {
    return &Client{
        client: openai.NewClient(apiKey),
    }
}

// GenerateCommitMessage generates a commit message from diff
func (c *Client) GenerateCommitMessage(diff string) (*models.CommitMessage, error) {
    prompt := fmt.Sprintf(`Generate a concise, conventional commit message for the following git diff. 
Follow conventional commit format (type(scope): description).

Diff:
%s

Return only a JSON object with this structure:
{
  "message": "commit message here"
}`, diff)

    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT4o,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: prompt,
                },
            },
            Temperature: 0.3,
            ResponseFormat: &openai.ChatCompletionResponseFormat{
                Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
                JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
                    Name: "CommitMessage",
                    Schema: json.RawMessage(`{
						"type": "object",
						"properties": {
							"message": {
								"type": "string"
							}
						},
						"required": ["message"]
					}`),
                },
            },
        },
    )

    if err != nil {
        return nil, fmt.Errorf("failed to generate commit message: %w", err)
    }

    var commitMsg models.CommitMessage
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &commitMsg); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    return &commitMsg, nil
}

// GeneratePRDetails generates PR title and description from diff
func (c *Client) GeneratePRDetails(diff string) (*models.PrDetails, error) {
    prompt := fmt.Sprintf(`Analyze the following git diff and generate a PR title, description, and file summaries.

Diff:
%s

Return a JSON object with this structure:
{
  "title": "PR title here",
  "description": "Detailed PR description here",
  "fileSummaries": [
    {
      "file": "path/to/file",
      "description": "One-sentence summary of changes"
    }
  ]
}`, diff)

    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT4,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: prompt,
                },
            },
            Temperature: 0.3,
            ResponseFormat: &openai.ChatCompletionResponseFormat{
                Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
                JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
                    Name: "PrDetails",
                    Schema: json.RawMessage(`{
						"type": "object",
						"properties": {
							"title": {
								"type": "string"
							},
							"description": {
								"type": "string"
							},
							"fileSummaries": {
								"type": "array",
								"items": {
									"type": "object",
									"properties": {
										"file": {
											"type": "string"
										},
										"description": {
											"type": "string"
										}
									},
									"required": ["file", "description"]
								}
							}
						},
						"required": ["title", "description", "fileSummaries"]
					}`),
                },
            },
        },
    )

    if err != nil {
        return nil, fmt.Errorf("failed to generate PR details: %w", err)
    }

    var prDetails models.PrDetails
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &prDetails); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    return &prDetails, nil
}

// ReviewPR generates review comments for a PR diff
func (c *Client) ReviewPR(diff string) (*models.PrReviewDetails, error) {
    prompt := fmt.Sprintf(`Review the following git diff and provide constructive feedback. 
Focus on code quality, security, performance, and best practices.

Diff:
%s

Return a JSON object with this structure:
{
  "review": [
    {
      "file": "path/to/file",
      "line": 42,
      "category": "Security|Bug|Optimization|Improvement",
      "comment": "Review comment here",
      "codeSnippet": "relevant code snippet"
    }
  ]
}`, diff)

    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: openai.GPT4oMini,
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: prompt,
                },
            },
            Temperature: 0.3,
						ResponseFormat: &openai.ChatCompletionResponseFormat{
							Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
							JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
								Name: "PrReviewDetails",
								Schema: json.RawMessage(`{
									"type": "object",
									"properties": {
										"review": {
											"type": "array",
											"items": {
												"type": "object",
												"properties": {
													"file": {"type": "string"},
													"line": {"type": "integer"},
													"category": {"type": "string"},
													"comment": {"type": "string"},
													"codeSnippet": {"type": "string"}
												},
												"required": ["file", "line", "category", "comment"]
											}
										}
									},
									"required": ["review"]
								}`),
							},
						},
        },
    )

    if err != nil {
        return nil, fmt.Errorf("failed to generate PR review: %w", err)
    }

    var reviewDetails models.PrReviewDetails
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &reviewDetails); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    return &reviewDetails, nil
}