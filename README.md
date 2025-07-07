# LangChain Go Gallery

A demonstration project showcasing AI integrations using OpenAI and Gemini AI APIs in Go.

## Features

- **OpenAI Integration**: Streaming chat completions using GPT-3.5-turbo
- **Gemini AI Integration**: 
  - Image description using Gemini-2.5-pro
  - Text generation using Gemini-2.5-flash

## Prerequisites

- Go 1.24.3 or higher
- OpenAI API key
- Gemini API key

## Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd langchaingo-gallery
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set environment variables**:
   ```bash
   export OPENAI_API_KEY="your-openai-api-key"
   export GEMINI_API_KEY="your-gemini-api-key"
   ```

## Build and Run

### Build the project:
```bash
go build -o langchaingo-gallery
```

### Run the executable:
```bash
./langchaingo-gallery
```

### Or run directly:
```bash
go run main.go
```

## What it does

The application demonstrates three AI functionalities:

1. **OpenAI Chat**: Asks GPT-3.5-turbo about Napoleon's defeat at Waterloo with streaming responses
2. **Gemini Image Description**: Analyzes and describes a sample image from the web
3. **Gemini Text Generation**: Generates a brief explanation of how AI works

## Dependencies

- `github.com/sashabaranov/go-openai` - OpenAI API client
- `github.com/google/generative-ai-go` - Gemini AI API client
- `google.golang.org/api` - Google API client libraries
- `google.golang.org/genai` - Additional Gemini AI support
