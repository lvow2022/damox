package web

type ChatRequest struct {
	Prompt string `json:"prompt"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
