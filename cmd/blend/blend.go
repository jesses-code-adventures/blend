package main

import (
	"fmt"
	"os"

	e "github.com/jesses-code-adventures/blend/env"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "blend",
	Short: "cli tool for generating git diffs from an LLM.",
}

func main() {
	e.LoadEnvVars()
	rootCmd.AddCommand(streamCmd)
	rootCmd.AddCommand(diffCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
