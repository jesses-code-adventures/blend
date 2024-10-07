package main

import (
	"context"
	"io"

	r "github.com/jesses-code-adventures/blend/runner"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "prompt the llm to create a diff between the current codebase and the desired outcome.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prompt := args[0]
		ctx := context.Background()
		staticSystemPrompt := `You are a developer tool for generating code.
You should respond in a specific format to every request. Every response should be in the form of a list of codeblocks.
Each set of codeblocks should be labelled after the first three backticks with the corresponding file path provided in later in the prompt.
You should only provide with codeblocks where you suggest diffs. If you have no diffs to suggest, you should return an empty set of backticks.
`
		runner := r.NewUnixChatGptRunner(ctx, staticSystemPrompt)
		reader, err := runner.RefreshRun(prompt)
		if err != nil {
			panic(err)
		}
		defer reader.Close()
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
			print(string(buf[:n]))
		}
	},
}
