package llm

import (
	"context"
	"fmt"
	"io"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AlreadyStreamingError struct{}

func (a AlreadyStreamingError) Error() string {
	return "There's already an open stream from OpenAi"
}

type NoApiKeyError struct{}

func (a NoApiKeyError) Error() string {
	return "Pass an api key ya dummy!"
}

type OpenAi struct {
	client       *openai.Client
	ctx          context.Context
	systemPrompt string
	streaming    bool
}

func NewOpenAi(ctx context.Context, openaiApiKey string) (OpenAi, error) {
	if openaiApiKey == "" {
		return OpenAi{}, NoApiKeyError{}
	}
	client := openai.NewClient(option.WithAPIKey(openaiApiKey))
	return OpenAi{
		client:       client,
		ctx:          ctx,
		systemPrompt: "",
		streaming:    false,
	}, nil
}

func (o *OpenAi) StopStreaming() {
	o.streaming = false
}

func (o *OpenAi) SetSystemPrompt(prompt string) {
	o.systemPrompt = prompt
}

func (o OpenAi) StreamTokens(prompt string) (pr io.ReadCloser, err error) {
	if o.streaming {
		return pr, AlreadyStreamingError{}
	}
	o.streaming = true
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		defer o.StopStreaming()
		stream := o.client.Chat.Completions.NewStreaming(o.ctx, openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(prompt),
				openai.SystemMessage(o.systemPrompt),
			}),
			Seed:  openai.Int(0),
			Model: openai.F(openai.ChatModelChatgpt4oLatest), // O1 doesn't work yet
		})
		acc := openai.ChatCompletionAccumulator{}
		for stream.Next() {
			chunk := stream.Current()
			acc.AddChunk(chunk)
			// TODO: Enable if using tool calls https://arc.net/l/quote/rlsqcacx
			// if tool, ok := acc.JustFinishedToolCall(); ok {
			// 	println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
			// }
			if refusal, ok := acc.JustFinishedRefusal(); ok {
				println("Refusal stream finished:", refusal)
			}
			if len(chunk.Choices) > 0 {
				content := []byte(chunk.Choices[0].Delta.Content)
				_, err := pw.Write(content)
				if err != nil {
					return
				}
			}
		}
	}()
	return pr, nil
}

func (o OpenAi) StreamPrint(prompt string) {
	r, err := o.StreamTokens(prompt)
	// TODO: Handle AlreadyStreamingError differently
	if err != nil {
		panic(err)
	}
	defer r.Close()
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(string(buf[:n]))
	}
}
