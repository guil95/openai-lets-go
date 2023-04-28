package entrypoints

import (
	"github.com/guil95/openai-lets-go/internal/chatgpt/entrypoints"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func HttpEntrypoint() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Start terminal",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go entrypoints.RunHTTP(quit)

			<-quit
		},
	}
}
