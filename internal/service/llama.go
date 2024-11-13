package service

import (
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"golang.org/x/net/context"
	"log"
)

type LlamaService interface {
	Chat(ctx context.Context, prompt string) (string, error)
}

type llamaService struct {
	llm *ollama.LLM
}

func NewLlamaService(llm *ollama.LLM) LlamaService {
	return &llamaService{
		llm: llm,
	}
}

func (l *llamaService) Chat(ctx context.Context, prompt string) (string, error) {
	completion, err := l.llm.Call(ctx, prompt,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return completion, nil
}
