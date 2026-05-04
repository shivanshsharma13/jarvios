package provider

import (
	"context"
	"strings"

	ollamaapi "github.com/ollama/ollama/api"
)

func OllamaChat(message string) (string, error) {
	client, err := ollamaapi.ClientFromEnvironment()
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	stream := true

	req := &ollamaapi.ChatRequest{
		Model:  "phi3",
		Stream: &stream,
		Messages: []ollamaapi.Message{
			{Role: "user", Content: message},
		},
	}

	err = client.Chat(context.Background(), req, func(resp ollamaapi.ChatResponse) error {
		sb.WriteString(resp.Message.Content)
		return nil
	})

	return sb.String(), err
}
