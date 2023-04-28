package client

type response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}
