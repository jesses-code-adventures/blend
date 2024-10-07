package llm

import (
	"io"
)

type Llm interface {
	SetSystemPrompt(prompt string)
	StreamTokens(input string) (io.ReadCloser, error)
	StreamPrint(input string)
}
