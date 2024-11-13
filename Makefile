# Makefile
.PHONY: wire build clean web

wire:
	@echo "Generating wire dependencies..."
	wire


build: wire
	@echo "Building the application..."
	go mod tidy
	go build  -v -o server ./

clean:
	@rm -f wire_gen.go

generate: web service repo
# 生成指定层的文件
web:
	go run ./cmd/generate/generate.go web $(name)

service:
	go run ./cmd/generate/generate.go service $(name)

repo:
	go run ./cmd/generate/generate.go repository $(name)

ddd:
	@echo "Creating DDD core directories..."
	# 创建核心目录
	mkdir -p internal/web internal/service internal/domain internal/repository
	mkdir -p ioc
	mkdir -p pkg
	mkdir -p config