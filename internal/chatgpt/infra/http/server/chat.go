package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/guil95/openai-lets-go/internal/chatgpt"
	"github.com/guil95/openai-lets-go/internal/chatgpt/commands"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ChatGPT echo.HandlerFunc

func HandleChat(service chatgpt.Chat, ctx context.Context) ChatGPT {
	return func(c echo.Context) error {
		var req ChatRequest

		if (&echo.DefaultBinder{}).BindHeaders(c, &req) != nil {
			return c.NoContent(http.StatusUnprocessableEntity)
		}

		if c.Bind(&req) != nil {
			return c.NoContent(http.StatusUnprocessableEntity)
		}

		answer, err := service.Completions(ctx, commands.Command{RequestID: req.RequestID, ApiKey: req.ApiKey, Text: req.Text})
		if err != nil {
			zap.S().Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		answer = regexp.MustCompile(`AI:`).ReplaceAllString(answer, "")
		answer = regexp.MustCompile(`Robô:`).ReplaceAllString(answer, "")

		return c.JSON(http.StatusOK, response{answer})
	}
}
