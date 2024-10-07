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
		runner := r.NewUnixChatGptRunner(ctx)
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
