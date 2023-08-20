package OpenAI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkoukk/tiktoken-go"
)

// Chat message roles defined by the OpenAI API.
const (
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleFunction  = "function"
)

type DataType string

const (
	Object  DataType = "object"
	Number  DataType = "number"
	Integer DataType = "integer"
	String  DataType = "string"
	Array   DataType = "array"
	Null    DataType = "null"
	Boolean DataType = "boolean"
)

// Definition is a struct for describing a JSON Schema.
// It is fairly limited, and you may have better luck using a third-party library.
type Definition struct {
	// Type specifies the data type of the schema.
	Type DataType `json:"type,omitempty"`
	// Description is the description of the schema.
	Description string `json:"description,omitempty"`
	// Enum is used to restrict a value to a fixed set of values. It must be an array with at least
	// one element, where each element is unique. You will probably only use this with strings.
	Enum []string `json:"enum,omitempty"`
	// Properties describes the properties of an object, if the schema type is Object.
	Properties map[string]Definition `json:"properties"`
	// Required specifies which properties are required, if the schema type is Object.
	Required []string `json:"required,omitempty"`
	// Items specifies which data type an array contains, if the schema type is Array.
	Items *Definition `json:"items,omitempty"`
}

func (d Definition) MarshalJSON() ([]byte, error) {
	if d.Properties == nil {
		d.Properties = make(map[string]Definition)
	}
	type Alias Definition
	return json.Marshal(struct {
		Alias
	}{
		Alias: (Alias)(d),
	})
}

type OpenAI struct {
	ApiKey string
}

func Create(apiKey string) *OpenAI {
	return &OpenAI{
		ApiKey: apiKey,
	}
}

func (openAI *OpenAI) CalculateTokens(request OpenAIChatRequest) (tokens int, err error) {
	return CalculateTokens(request)
}

func CalculateTokens(request OpenAIChatRequest) (tokens int, err error) {
	tke, err := tiktoken.GetEncoding("cl100k_base")
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return
	}

	for _, message := range request.Messages {
		tokens += len(tke.Encode(message.Content, nil, nil))
	}

	return
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
