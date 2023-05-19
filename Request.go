package OpenAI

type OpenAIChatRequest struct {
	Model       string    `json:"model,omitempty"`
	Messages    []Message `json:"messages,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
	N           int       `json:"n,omitempty"`
	//? Disabled for now...
	// Stream           bool                `json:"stream,omitempty"`
	Stop             StringOrStringArray `json:"stop,omitempty"`
	MaxTokens        int                 `json:"max_tokens,omitempty"`
	PresencePenalty  float64             `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64             `json:"frequency_penalty,omitempty"`
	LogitBias        map[uint]int        `json:"logit_bias,omitempty"`
	User             string              `json:"user,omitempty"`
}

const API_URL = "https://api.openai.com/v1/chat/completions"
