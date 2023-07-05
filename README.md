# A simple wrapper for the [OpenAI REST API](https://platform.openai.com/docs/api-reference/chat/create) written in pure Golang.

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/cr2007/openai-go)

[![Go Reference](https://pkg.go.dev/badge/github.com/Gaurav-Gosain/openai-go.svg)](https://pkg.go.dev/github.com/Gaurav-Gosain/openai-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Gaurav-Gosain/openai-go)](https://goreportcard.com/report/github.com/Gaurav-Gosain/openai-go)


> Note: Currently only supports the Chat API (without streaming).

### Supported Models

|   Model Name    |
| :-------------: |
|      gpt-4      |
|     gpt-4-*     |
|  gpt-3.5-turbo  |
| gpt-3.5-turbo-* |


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
	ai := OpenAI.Create("YOUR_API_KEY_HERE") // Ideally, you should store your API key in an environment variable

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

	response, err := ai.Chat(chatRequest)

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

## Calculate Request Tokens without calling the OpenAI API

Uses the [tiktoken-go](https://github.com/pkoukk/tiktoken-go) and the c100k_base encoding model to estimate the number of tokens in a request. 

> Note: When making a chat request to the OpenAI APIs, the System Prompt, the user messages and the assistant responses all count towards the total tokens.

> Credits: [tiktoken-go](https://github.com/pkoukk/tiktoken-go)

```go
package main

import (
	"fmt"

	OpenAI "github.com/Gaurav-Gosain/openai-go"
)

func main() {
	//? Get your API key from https://platform.openai.com/account/api-keys
	ai := OpenAI.Create("YOUR_API_KEY_HERE") // Ideally, you should store your API key in an environment variable

	//? Refer to https://platform.openai.com/docs/api-reference/chat/create to learn more about the parameters.
	chatRequest := OpenAI.OpenAIChatRequest{
		Model:     "gpt-3.5-turbo", // name of the model to use
		MaxTokens: 50,              // max tokens to generate
		Messages: []OpenAI.Message{ // messages to feed to the model
			{
				Role:    OpenAI.ChatMessageRoleSystem,
				Content: "This is an example of a system prompt.",
			},
			{
				Role:    OpenAI.ChatMessageRoleUser,
				Content: "Hello World! This is an example of a chat message.",
			},
			{
				Role:    OpenAI.ChatMessageRoleAssistant,
				Content: "Hello World! This is an example of an assistant response.",
			},
		},
	}

	//? Calculate the number of tokens that will be consumed by the request.
	tokens, err := ai.CalculateTokens(chatRequest)

	//? OR
	// tokens, err := OpenAI.CalculateTokens(chatRequest)

	if err != nil {
		fmt.Println("ERROR:", err)
		//! Handle Tokenizer error
		return
	}

	fmt.Println("This request has", tokens, "tokens")
    // Output: This request has 33 tokens
}
```

Work In progress...

This module is not complete and does not support all features offered by the OpenAI API currently. Any contibutions are welcome!

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
