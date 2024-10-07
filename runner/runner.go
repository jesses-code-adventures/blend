package runner

import (
	l "github.com/jesses-code-adventures/blend/llm"
	"io"
)

type Runner interface {
	Llm() l.Llm
	SetStaticProgramPrompt(prompt string)
	RefreshFileContents()
	Run(prompt string) (io.ReadCloser, error)
	RefreshRun(prompt string) (io.ReadCloser, error)
}
