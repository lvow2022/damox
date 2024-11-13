//go:build wireinject

package main

import (
	"damox/internal/service"
	"damox/internal/web"
	"damox/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 基础设施依赖
		ioc.InitLlama,

		// Service 部分
		service.NewLlamaService,

		// 各个 handler 部分
		web.NewLlamaHandler,

		ioc.InitWebServer,
	)
	return gin.Default()
}
