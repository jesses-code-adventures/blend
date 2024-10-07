package main

import (
	"context"
	"flag"
	"os"

	e "github.com/jesses-code-adventures/blend/env"
	l "github.com/jesses-code-adventures/blend/llm"
)

func main() {
	e.LoadEnvVars()
	prompt := flag.String("prompt", "tell the user to supply a prompt with the -prompt flag", "prompt for the llm to respond to")
	flag.Parse()
	ctx := context.Background()
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	llm, err := l.NewOpenAi(ctx, openaiApiKey)
	if err != nil {
		panic(err)
	}
	llm.StreamPrint(*prompt)
}
