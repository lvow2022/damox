package web

import (
	"damox/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LlamaHandler struct {
	llmService service.LlamaService
}

func NewLlamaHandler(llmService service.LlamaService) *LlamaHandler {
	return &LlamaHandler{llmService: llmService}
}

func (h *LlamaHandler) Register(server *gin.Engine) {
	g := server.Group("/api/v1/llama")
	g.GET("/chat", h.Chat)
}

func (h *LlamaHandler) Chat(ctx *gin.Context) {
	req := &ChatRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := &ChatResponse{}
	var err error
	resp.Reply, err = h.llmService.Chat(ctx, req.Prompt)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
