package OpenAI

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//? Get your API key from https://platform.openai.com/account/api-keys
	OpenAI := Create("YOUR_API_KEY_HERE")

	//? Refer to https://platform.openai.com/docs/api-reference/chat/create to learn more about the parameters.
	chatRequest := OpenAIChatRequest{
		Model:     "gpt-3.5-turbo", // name of the model to use
		MaxTokens: 50,              // max tokens to generate
		Messages: []Message{ // messages to feed to the model
			{
				Role:    ChatMessageRoleUser,
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

	response, err := OpenAI.Chat(chatRequest)

	if err != nil {
		t.Fail()
	}

	if response.Error.Type != "" {

		fmt.Println("ERROR:", response.Error.Type, ":", response.Error.Code)

		t.Fail()
		return
	}

	for _, choice := range response.Choices {
		fmt.Printf("%s: %s\n", choice.Message.Role, choice.Message.Content)
	}

	fmt.Println("That costed you", response.Usage.TotalTokens, "tokens")

}
