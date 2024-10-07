package llm

import (
	"context"
	"os"
	"testing"

	e "github.com/jesses-code-adventures/blend/env"
)

func Test_StopStreaming(t *testing.T) {
	e.LoadEnvVarsWithTestVars()
	ctx := context.Background()
	openai, err := NewOpenAi(ctx, os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		t.Fatalf("failed to construct opeani with key %s", os.Getenv("OPENAI_API_KEY"))
	}
	openai.streaming = true
	openai.StopStreaming()
	if openai.streaming == true {
		t.Errorf("streaming was not stopped properly")
	}
}
