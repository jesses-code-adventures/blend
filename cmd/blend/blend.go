package main

import (
	"context"
	"fmt"
	"os"

	e "github.com/jesses-code-adventures/blend/env"
	l "github.com/jesses-code-adventures/blend/llm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blend",
	Short: "cli tool for generating git diffs from an LLM.",
}

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

func main() {
	e.LoadEnvVars()
	rootCmd.AddCommand(streamCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
