package main

import (
	"github.com/guil95/openai-lets-go/cmd/entrypoints"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "split-entry-points",
	}

	rootCmd.AddCommand(entrypoints.HttpEntrypoint())
	rootCmd.AddCommand(entrypoints.TerminalEntrypoint())
	_ = rootCmd.Execute()
}

/*
func main() {
	e := echo.New()
	ctx := context.Background()
	logger.SetupLogger(ctx)

	service := chatgpt.NewService(client.NewChatGptClient(), repository.NewCommandMongoRepository(storages.Connect(ctx)))
	e.POST("/chat", echo.HandlerFunc(server.HandleChat(service, ctx)))

	e.Logger.Fatal(e.Start(":8989"))
}
*/

//}
//func main() {
//	apiKey := os.Getenv("OPENAI_API_KEY")
//	text := ""
//	ctx := context.Background()
//	for {
//		service := chatgpt.NewService(client.NewChatGptClient(), repository.NewCommandMongoRepository(storages.Connect(ctx)))
//
//		fmt.Println("What is your question?")
//		in := bufio.NewReader(os.Stdin)
//		scanText, err := in.ReadString('\n')
//		if err != nil {
//			panic(err)
//		}
//
//		text = fmt.Sprintf("%sHuman:%s", text, scanText)
//
//		answer, err := service.Completions(ctx, commands.Command{ApiKey: apiKey, Text: text})
//		if err != nil {
//			panic(err)
//		}
//		text = fmt.Sprintf("%sAI:%s", text, answer)
//
//		zap.S().Info(answer)
//	}
//
//}
