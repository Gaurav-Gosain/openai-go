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
}
