package main

import (
	"fmt"

	OpenAI "github.com/Gaurav-Gosain/openai-go"
)

func main() {
	//? Get your API key from https://platform.openai.com/account/api-keys
	ai := OpenAI.Create("OPEN_API_KEY") // Ideally, you should store your API key in an environment variable

	jsonSchema := OpenAI.Definition{
		Type: OpenAI.Object,
		Properties: map[string]OpenAI.Definition{
			"Name": {
				Type:        OpenAI.String,
				Description: "The name of the person.",
				Enum: []string{
					"John Doe", "Jane Doe", "John Smith", "Jane Smith",
				},
			},
			"Age": {Type: OpenAI.Number},
		},
		Required: []string{"Name", "Age"},
	}

	//? Refer to https://platform.openai.com/docs/api-reference/chat/create to learn more about the parameters.
	chatRequest := OpenAI.OpenAIChatRequest{
		Model:     OpenAI.GPT3Dot5Turbo0613, // name of the model to use
		MaxTokens: 50,                       // max tokens to generate
		Messages: []OpenAI.Message{ // messages to feed to the model
			{
				Role:    OpenAI.ChatMessageRoleUser,
				Content: "Hi my name is John Doe and I am 25 years old.",
			},
		},
		Functions: []OpenAI.FunctionDefinition{
			{
				Name:       "valid_json",
				Parameters: jsonSchema,
			},
		},
		FunctionCall: OpenAI.FunctionCall{
			Name: "valid_json",
		},
		Stop:             []string{"\n"}, // strings and string arrays are both accepted (upto 4 stop tokens)
		Temperature:      1,
		TopP:             1,
		N:                1,
		PresencePenalty:  0.6,
		FrequencyPenalty: 0.6,
		LogitBias: map[string]int{ // https://platform.openai.com/tokenizer
			"2188":  100,
			"42392": 100,
		},
		User: "user_123",
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

	response, err := ai.Chat(chatRequest)

	if err != nil {
		fmt.Println("ERROR:", err)
		//! Handle REST API error
		return
	}

	if response.Error.Type != "" {
		fmt.Println("ERROR:", response.Error.Type, ":", response.Error.Code)
		//! Handle OpenAPI Response error
		return
	}

	for _, choice := range response.Choices {
		if choice.Message.Content != "" {
			fmt.Printf("%s: %s\n", choice.Message.Role, choice.Message.Content)
		}
		if choice.Message.FunctionCall != nil {
			fmt.Printf("Function Call\n%+v\n", choice.Message.FunctionCall)
		}
	}

	fmt.Println("That costed you", response.Usage.TotalTokens, "tokens")
}
