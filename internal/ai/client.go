package ai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/richardamare/gitai/internal/models"
	"github.com/sashabaranov/go-openai"
)

const model = openai.GPT4oMini

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
    prompt := fmt.Sprintf(commitMessagePrompt, diff)

    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: model,
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

// GenerateMRDetails generates MR title and description from diff
func (c *Client) GenerateMRDetails(diff string) (*models.MrDetails, error) {
    prompt := fmt.Sprintf(`Analyze the following git diff and generate a MR title, description, and file summaries.

Diff:
%s

Return a JSON object with this structure:
{
  "title": "MR title here",
  "description": "Detailed MR description here",
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
            Model: model,
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

    var prDetails models.MrDetails
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &prDetails); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    return &prDetails, nil
}

// GenerateMRTitle generates a concise PR title from a diff
func (c *Client) GenerateMRTitle(diff string) (string, error) {
	prompt := fmt.Sprintf(mrTitlePrompt, diff)

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
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
					Name: "PrTitle",
					Schema: json.RawMessage(`{
						"type": "object",
						"properties": {
							"title": {
								"type": "string"
							}
						},
						"required": ["title"]
					}`),
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate PR title: %w", err)
	}

	var prTitle models.MrTitle
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &prTitle); err != nil {
		return "", fmt.Errorf("failed to parse AI response: %w", err)
	}

	return prTitle.Title, nil
}

// ReviewMR generates review comments for a MR diff
func (c *Client) ReviewMR(diff string) (*models.MrReviewDetails, error) {
    prompt := fmt.Sprintf(reviewPrompt, diff)

    resp, err := c.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: model,
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

    var reviewDetails models.MrReviewDetails
    if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &reviewDetails); err != nil {
        return nil, fmt.Errorf("failed to parse AI response: %w", err)
    }

    return &reviewDetails, nil
}