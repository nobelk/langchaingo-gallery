package openai

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

func RunLLM() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Why did Napoleon Bonaparte loose the Battle of Waterloo?",
			},
		},
		Stream: true,
	}

	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	for {
		msg, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(msg.Choices[0].Delta.Content)
	}
}
