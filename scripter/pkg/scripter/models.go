package scripter

import (
	"net/url"
	"strconv"
	"strings"
)

type RequestMessage struct {
	Role    string `json:"role"` // "system", "user", or "assistant"
	Content string `json:"content"`
}

type RequestMessages []RequestMessage

type ModelInfo struct {
	Provider    string
	Model       string
	Host        string
	QueryParams map[string]float64
}

// Build user message
func BuildMessage(content string) RequestMessage {
	return RequestMessage{
		Role:    "user",
		Content: content,
	}
}

func GetModelAndProvider(model string) (ModelInfo, error) {
	parsed, err := url.Parse(model)
	if err != nil {
		return ModelInfo{}, err
	}

	provider := parsed.Scheme
	host := ""
	modelName := parsed.Host

	if strings.Contains(parsed.Host, "@") {
		parts := strings.Split(parsed.Host, "@")
		modelName, host = parts[0], parts[1]
	} else if provider == "openrouter" {
		modelName = parsed.Host + parsed.Path
	}

	params := make(map[string]float64)
	for key, values := range parsed.Query() {
		if val, err := strconv.ParseFloat(values[0], 64); err == nil {
			params[key] = val
		}
	}
	if len(params) == 0 {
		params = nil
	}

	return ModelInfo{
		Provider:    provider,
		Model:       modelName,
		Host:        host,
		QueryParams: params,
	}, nil
}

func (messages RequestMessages) GetLastMessage() string {
	if len(messages) == 0 {
		return ""
	}
	return messages[len(messages)-1].Content
}
