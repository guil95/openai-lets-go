package entrypoints

import (
	"context"
	"github.com/guil95/openai-lets-go/config/logger"
	"github.com/guil95/openai-lets-go/config/storages"
	"github.com/guil95/openai-lets-go/internal/chatgpt"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/http/client"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/http/server"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/repository"
	"github.com/labstack/echo/v4"
	"os"
)

func RunHTTP(quit chan os.Signal) {
	e := echo.New()
	ctx := context.Background()
	logger.SetupLogger(ctx)

	service := chatgpt.NewService(client.NewChatGptClient(), repository.NewCommandMongoRepository(storages.Connect(ctx)))
	e.POST("/chat", echo.HandlerFunc(server.HandleChat(service, ctx)))

	e.Logger.Fatal(e.Start(":8989"))
}
