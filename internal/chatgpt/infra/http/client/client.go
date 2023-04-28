package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ChatGptClient interface {
	Completions(text, apiKey string) (string, error)
}

const openaiUrl = "https://api.openai.com/v1/completions"

type chatGptClient struct{}

func NewChatGptClient() ChatGptClient {
	return &chatGptClient{}
}

func (c chatGptClient) Completions(text, apiKey string) (string, error) {
	reqBytes, err := json.Marshal(c.buildPayload(text))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req, err := http.NewRequest("POST", openaiUrl, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	respText, err := c.retrieveResponse(resp)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(respText)
	}

	return respText, nil
}

func (c chatGptClient) buildPayload(text string) completionsPayload {
	return completionsPayload{
		Model:            "text-davinci-003",
		Prompt:           text,
		Temperature:      0.9,
		MaxTokens:        2000,
		TopP:             1,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.6,
		Stop:             []string{" Human:", " AI:"},
	}
}

func (c chatGptClient) retrieveResponse(httpResponse *http.Response) (string, error) {
	if httpResponse.StatusCode == http.StatusOK {
		var r response
		err := json.NewDecoder(httpResponse.Body).Decode(&r)
		if err != nil {
			fmt.Println(err)
			return "", err
		}

		return r.Choices[0].Text, nil
	}

	var respBody map[string]interface{}
	err := json.NewDecoder(httpResponse.Body).Decode(&respBody)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return respBody["error"].(map[string]interface{})["message"].(string), nil
}
