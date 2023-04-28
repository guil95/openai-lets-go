package entrypoints

import (
	"github.com/guil95/openai-lets-go/internal/chatgpt/entrypoints"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func TerminalEntrypoint() *cobra.Command {
	return &cobra.Command{
		Use:   "terminal",
		Short: "Start terminal",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go entrypoints.RunTerminal(quit)
			<-quit
		},
	}
}
