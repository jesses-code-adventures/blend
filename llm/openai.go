package llm

import (
	"context"
	"io"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type AlreadyStreamingError struct{}

func (a AlreadyStreamingError) Error() string {
	return "There's already an open stream from OpenAi"
}

type OpenAi struct {
	client       *openai.Client
	systemPrompt string
	streaming    bool
}

func NewOpenAi(openaiApiKey string) OpenAi {
	return OpenAi{
		client:       openai.NewClient(option.WithAPIKey(openaiApiKey)),
		systemPrompt: "",
		streaming:    false,
	}
}

func (o *OpenAi) SetSystemPrompt(prompt string) {
	o.systemPrompt = prompt
}

func (o *OpenAi) StopStreaming() {
	o.streaming = false
}

func (o *OpenAi) StreamTokens(ctx context.Context, prompt string) (pr io.ReadCloser, err error) {
	if o.streaming {
		return pr, AlreadyStreamingError{}
	}
	o.streaming = true
	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		defer o.StopStreaming()
		stream := o.client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
			Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(o.systemPrompt),
				openai.UserMessage(prompt),
			}),
			Seed:  openai.Int(0),
			Model: openai.F(openai.ChatModelO1Mini),
		})

		acc := openai.ChatCompletionAccumulator{}
		for stream.Next() {
			chunk := stream.Current()
			acc.AddChunk(chunk)

			if content, ok := acc.JustFinishedContent(); ok {
				println("Content stream finished:", content)
			}

			// TODO: Enable if using tool calls https://arc.net/l/quote/rlsqcacx
			// if tool, ok := acc.JustFinishedToolCall(); ok {
			// 	println("Tool call stream finished:", tool.Index, tool.Name, tool.Arguments)
			// }

			if refusal, ok := acc.JustFinishedRefusal(); ok {
				println("Refusal stream finished:", refusal)
			}

			if len(chunk.Choices) > 0 {
				_, err := pw.Write([]byte(chunk.Choices[0].Delta.Content))
				if err != nil {
					return
				}
			}
		}
	}()
	return pr, nil
}
