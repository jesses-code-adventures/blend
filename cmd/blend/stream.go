package main

import (
	"context"
	l "github.com/jesses-code-adventures/blend/llm"
	"github.com/spf13/cobra"
	"os"
)

var streamCmd = &cobra.Command{
	Use:   "stream [prompt]",
	Short: "stream tokens straight from the LLM and print them.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		ctx := context.Background()
		openaiApiKey := os.Getenv("OPENAI_API_KEY")
		llm, err := l.NewOpenAi(ctx, openaiApiKey)
		if err != nil {
			panic(err)
		}
		llm.StreamPrint(prompt)
	},
}
