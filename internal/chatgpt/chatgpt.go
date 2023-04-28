package chatgpt

import (
	"context"
	"fmt"
	"github.com/guil95/openai-lets-go/internal/chatgpt/commands"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/http/client"
	"go.uber.org/zap"
)

type Chat interface {
	Completions(ctx context.Context, command commands.Command) (string, error)
}

type service struct {
	openai client.ChatGptClient
	repo   Repository
}

func NewService(openai client.ChatGptClient, repo Repository) Chat {
	return &service{openai, repo}
}

func (s service) Completions(ctx context.Context, command commands.Command) (string, error) {
	chat, err := s.repo.GetChatByEntityID(ctx, command.RequestID)
	if err != nil {
		zap.S().Error(err)
		return "", err
	}

	text := ""

	if chat != nil {
		text = fmt.Sprintf("%sHuman:%s", chat.Text, command.Text)
		completions, err := s.openai.Completions(text, command.ApiKey)
		if err != nil {
			return "", err
		}

		err = s.repo.UpdateChat(ctx, ChatGPT{Text: fmt.Sprintf("%sAI:%s", text, completions), EntityID: command.RequestID})
		if err != nil {
			zap.S().Error(fmt.Sprintf("error to update chat %v", err))
		}

		return completions, nil
	}

	text = fmt.Sprintf("Human:%s", command.Text)

	completions, err := s.openai.Completions(text, command.ApiKey)
	if err != nil {
		return "", err
	}

	go func() {
		err := s.repo.SaveChat(ctx, ChatGPT{Text: fmt.Sprintf("%sAI:%s", text, completions), EntityID: command.RequestID})
		if err != nil {
			zap.S().Error(fmt.Sprintf("error to save chat %v", err))
		}
	}()

	return completions, err
}

type ChatGPT struct {
	EntityID string `json:"entity_id" bson:"entity_id"`
	Text     string `json:"text" bson:"text"`
}
