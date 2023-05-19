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
		// Stop:        "\n", // strings and string arrays are both accepted (upto 4 stop tokens)
		// Temperature: 1,
		// TopP:        1,
		// N:           1,
		// PresencePenalty:  0.6,
		// FrequencyPenalty: 0.6,
		// LogitBias: map[uint]int{ // https://platform.openai.com/tokenizer
		// 	2188:  100,
		// 	42392: 100,
		// },
		// User: "user_123",
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
