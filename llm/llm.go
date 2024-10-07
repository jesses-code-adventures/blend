package llm

import (
	"context"
	"io"
)

type Llm interface {
	SetSystemPrompt(prompt string)
	StreamTokens(ctx context.Context, input string) (io.ReadCloser, error)
}
