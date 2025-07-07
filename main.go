package main

import (
	"fmt"
	"github.com/nobelk/langchaingo-gallery/geminiai"
	"github.com/nobelk/langchaingo-gallery/openai"
)

func main() {
	fmt.Println("\n====OpenAI====\n")
	openai.RunLLM()
	fmt.Println("\n ====GeminiAI====\n ")
	geminiai.DescribeImage()
	fmt.Println("\n")
	geminiai.GenerateText()
}
