package ioc

import (
	"damox/internal/web"
	"github.com/gin-gonic/gin"
)

func InitWebServer(LlamaHandler *web.LlamaHandler) *gin.Engine {
	server := gin.Default()
	LlamaHandler.Register(server)
	return server
}
