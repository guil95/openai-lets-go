package chatgpt

import "context"

type Repository interface {
	SaveChat(ctx context.Context, gpt ChatGPT) error
	UpdateChat(ctx context.Context, gpt ChatGPT) error
	GetChatByEntityID(ctx context.Context, entityID string) (*ChatGPT, error)
}
