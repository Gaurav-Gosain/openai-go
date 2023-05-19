# A simple wrapper for the [OpenAI REST API](https://platform.openai.com/docs/api-reference/chat/create) written in pure Golang.

> Note: Currently only supports the Chat API (without streaming).

## Installation

```bash
go get github.com/Gaurav-Gosain/openai-go
```

## Simple OpenAI Chat Rest API Example Usage

```go
package main

import (
	"fmt"

	OpenAI "github.com/Gaurav-Gosain/openai-go"
)

func main() {
	//? Get your API key from https://platform.openai.com/account/api-keys
	openAI := OpenAI.Create("YOUR_API_KEY_HERE") // Ideally, you should store your API key in an environment variable

	//? Refer to https://platform.openai.com/docs/api-reference/chat/create to learn more about the parameters.
	chatRequest := OpenAI.OpenAIChatRequest{
		Model:     "gpt-3.5-turbo", // name of the model to use
		MaxTokens: 50,              // max tokens to generate
		Messages: []OpenAI.Message{ // messages to feed to the model
			{
				Role:    OpenAI.ChatMessageRoleUser,
				Content: "Hello",
			},
		},
	}

	response, err := openAI.Chat(chatRequest)

	if err != nil {
		fmt.Println("ERROR:", err)
		// Handle REST API error
		return
	}

	if response.Error.Type != "" {
		fmt.Println("ERROR:", response.Error.Type, ":", response.Error.Code)
		// Handle OpenAPI Response error
		return
	}

	for _, choice := range response.Choices {
		fmt.Printf("%s: %s\n", choice.Message.Role, choice.Message.Content)
	}

	fmt.Println("That costed you", response.Usage.TotalTokens, "tokens")
}
```
