package ioc

import (
	"github.com/tmc/langchaingo/llms/ollama"
)

func InitLlama() *ollama.LLM {
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		panic(err)
	}
	return llm
}
