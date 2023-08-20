package main

import (
	"fmt"

	OpenAI "github.com/Gaurav-Gosain/openai-go"
)

func main() {
	//? Get your API key from https://platform.openai.com/account/api-keys
	ai := OpenAI.Create("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") // Ideally, you should store your API key in an environment variable

	jsonSchema := OpenAI.Definition{
		Type: OpenAI.Object,
		Properties: map[string]OpenAI.Definition{
			"Instrument": {
				Type: OpenAI.String,
				Enum: []string{
					"EURUSD", "EURAUD", "NZDJPY", "GBPUSD", "GBPJPY", "EURJPY", "USDJPY", "USDCHF", "EURCHF",
					"AUDUSD", "USDCAD", "NZDUSD", "GBPCHF", "AUDJPY", "AUDCAD", "AUDCHF", "AUDNZD", "CADCHF",
					"CHFJPY", "EURNZD", "EURCAD", "EURGBP", "EURSEK", "GBPAUD", "GBPCAD", "GBPNZD", "NZDCAD",
					"NZDCHF", "NZDJPY", "USDSGD", "USDTRY", "USDZAR", "XAUUSD", "XAGUSD",
				},
			},
			"OrderType": {
				Type: OpenAI.String,
				Enum: []string{
					"buy", "Buy", "BUY",
					"sell", "Sell", "SELL",
					"buy limit", "Buy Limit", "BUY LIMIT",
					"sell limit", "Sell Limit", "SELL LIMIT",
					"buy stop", "Buy Stop", "BUY STOP",
					"sell stop", "Sell Stop", "SELL STOP",
				},
			},
			"EntryPrice": {Type: OpenAI.Number},
			"SL":         {Type: OpenAI.Number},
			"TPs": {
				Type:  OpenAI.Array,
				Items: &OpenAI.Definition{Type: OpenAI.Number},
			},
		},
		Required: []string{"Instrument", "OrderType", "EntryPrice", "SL", "TPs"},
	}

	//? Refer to https://platform.openai.com/docs/api-reference/chat/create to learn more about the parameters.
	chatRequest := OpenAI.OpenAIChatRequest{
		Model:     OpenAI.GPT3Dot5Turbo0613, // name of the model to use
		MaxTokens: 50,                       // max tokens to generate
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
		Functions: []OpenAI.FunctionDefinition{
			{
				Name:       "valid_json",
				Parameters: jsonSchema,
			},
		},
		FunctionCall: OpenAI.FunctionCall{
			Name: "valid_json",
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
