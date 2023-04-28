package client

type completionsPayload struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      float64  `json:"temperature"`
	MaxTokens        int64    `json:"max_tokens"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}
