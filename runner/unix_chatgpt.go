package runner

import (
	"context"
	"fmt"
	"io"
	"os"

	i "github.com/jesses-code-adventures/blend/ingest"
	l "github.com/jesses-code-adventures/blend/llm"
)

type UnixChatGptRunner struct {
	ingestor            i.Ingestor
	llm                 l.Llm
	staticProgramPrompt string
}

func NewUnixChatGptRunner(ctx context.Context, staticSystemPrompt string) UnixChatGptRunner {
	llm, err := l.NewOpenAi(ctx, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		panic(err)
	}
	ingestor := i.NewUnixFilepathIngestor()
	return UnixChatGptRunner{ingestor: ingestor, llm: llm, staticProgramPrompt: staticSystemPrompt}
}

func (r *UnixChatGptRunner) Llm() l.Llm {
	return r.llm
}

func (r *UnixChatGptRunner) SetStaticProgramPrompt(prompt string) {
	r.staticProgramPrompt = prompt
}

func (r *UnixChatGptRunner) RefreshFileContents() {
	r.ingestor.Ingest()
}

func (r *UnixChatGptRunner) Run(prompt string) (io.ReadCloser, error) {
	return r.llm.StreamTokens(prompt)
}

func (r *UnixChatGptRunner) RefreshRun(prompt string) (io.ReadCloser, error) {
	r.RefreshFileContents()
	fileContents := r.ingestor.ContentsString()
	systemPrompt := fmt.Sprintf("***** PROGRAM DEFINITION *****\n\n%s\n\n***** FILE CONTENTS ***** \n\n%s", r.staticProgramPrompt, fileContents)
	r.llm.SetSystemPrompt(systemPrompt)
	return r.Run(prompt)
}
