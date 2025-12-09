package scripter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ClientConfig struct {
	authToken string
	BaseURL   string
}

type Client struct {
	config ClientConfig
}

func NewClient(auth string) *Client {
	config := DefaultConfig(auth)
	return &Client{*config}
}

func DefaultConfig(authToken string) *ClientConfig {
	return &ClientConfig{
		authToken: authToken,
		BaseURL:   "https://openrouter.ai/api/v1/chat/completions",
	}
}

func (c *Client) GenerateText(messages []RequestMessage, model string) ([]RequestMessage, error) {
	body, err := json.Marshal(OpenRouterRequest{
		Model:    model,
		Messages: messages,
	})
	if err != nil {
		return nil, fmt.Errorf("error formatting request body: %w", err)
	}

	request, err := http.NewRequest("POST", c.config.BaseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.config.authToken)
	request.Header.Set("Content-Type", "application/json")

	cl := &http.Client{}
	response, err := cl.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	output, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response body: %w", err)
	}

	var data OpenRouterResponse
	err = json.Unmarshal(output, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(data.Choices) == 0 {
		log.Printf("API returned invalid response. Full response: %s", string(output))
		return nil, fmt.Errorf("API returned no choices")
	}

	messages = append(messages, RequestMessage{
		Role:    data.Choices[0].Message.Role,
		Content: data.Choices[0].Message.Content,
	})

	return messages, nil
}
