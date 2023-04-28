package server

type ChatRequest struct {
	ApiKey    string `header:"x-api-key"`
	RequestID string `json:"request_id"`
	Text      string `json:"text"`
}
