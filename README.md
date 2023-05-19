# A simple wrapper for the [OpenAI REST API](https://platform.openai.com/docs/api-reference/chat/create) written in pure Golang.

[![Go Reference](https://pkg.go.dev/badge/github.com/Gaurav-Gosain/openai-go.svg)](https://pkg.go.dev/github.com/Gaurav-Gosain/openai-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Gaurav-Gosain/openai-go)](https://goreportcard.com/report/github.com/Gaurav-Gosain/openai-go)


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

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=Gaurav-Gosain/openai-go&type=Date)](https://star-history.com/#Gaurav-Gosain/openai-go&Date)

<div style="display:flex;flex-wrap:wrap;">
  <img alt="GitHub Language Count" src="https://img.shields.io/github/languages/count/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Top Language" src="https://img.shields.io/github/languages/top/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="" src="https://img.shields.io/github/repo-size/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Issues" src="https://img.shields.io/github/issues/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Closed Issues" src="https://img.shields.io/github/issues-closed/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Pull Requests" src="https://img.shields.io/github/issues-pr/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Closed Pull Requests" src="https://img.shields.io/github/issues-pr-closed/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Contributors" src="https://img.shields.io/github/contributors/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Last Commit" src="https://img.shields.io/github/last-commit/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
  <img alt="GitHub Commit Activity (Week)" src="https://img.shields.io/github/commit-activity/w/Gaurav-Gosain/openai-go" style="padding:5px;margin:5px;" />
<div>
