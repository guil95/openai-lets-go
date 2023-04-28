package entrypoints

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guil95/openai-lets-go/config/storages"
	"github.com/guil95/openai-lets-go/internal/chatgpt"
	"github.com/guil95/openai-lets-go/internal/chatgpt/commands"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/http/client"
	"github.com/guil95/openai-lets-go/internal/chatgpt/infra/repository"
	"os"
	"os/signal"
	"regexp"
)

func RunTerminal(quit chan os.Signal) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()

	fmt.Print("ChatID: ")
	in := bufio.NewReader(os.Stdin)
	entityID, err := in.ReadString('\n')
	if err != nil {
		signal.Notify(quit)

		panic(err)
	}

	for {
		service := chatgpt.NewService(client.NewChatGptClient(), repository.NewCommandMongoRepository(storages.Connect(ctx)))

		fmt.Println("What is your question?")
		in := bufio.NewReader(os.Stdin)
		text, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		answer, err := service.Completions(ctx, commands.Command{RequestID: entityID, ApiKey: apiKey, Text: text})
		if err != nil {
			panic(err)
		}

		answer = regexp.MustCompile(`AI:`).ReplaceAllString(answer, "")
		answer = regexp.MustCompile(`Robô:`).ReplaceAllString(answer, "")

		fmt.Println(answer)
	}
}
