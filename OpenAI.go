package OpenAI

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// Chat message roles defined by the OpenAI API.
const (
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleSystem    = "system"
)

type OpenAI struct {
	ApiKey string
}

func Create(apiKey string) *OpenAI {
	return &OpenAI{
		ApiKey: apiKey,
	}
}

func (openAI *OpenAI) Chat(request OpenAIChatRequest) (response OpenAIChatResponse, err error) {

	// Convert the struct to JSON.
	jsonData, err := json.Marshal(request)
	if err != nil {
		return
	}

	// Convert the JSON data to a string reader.
	payload := strings.NewReader(string(jsonData))

	req, _ := http.NewRequest("POST", API_URL, payload)

	req.Header.Add("Authorization", "Bearer "+openAI.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}

	json.Unmarshal(body, &response)

	return response, nil
}
