# Makefile for landers1037/blog
# 2021.7.23
# 使用默认方式动态链接构建

# envs
APP_NAME=blogo
BUILD_STATIC=-extldflags -static
GO111MODULES=on
CGO_ENABLED=1
GOPROXY=https://goproxy.io,direct

.PHONY: build
## build: build the blog application
build: clean env
	@echo "start to build app blog"
	go build --mod=mod -a -ldflags="-w -s ${BUILD_STATIC}" -trimpath -tags="nomsgpack go_json" -o ${APP_NAME}

.PHONY: run
## run: run go run app.go with -race
run:
	@go run -race app.go

.PHONY: clean
## clean: clean the go binary
clean:
	@echo "start to clean the binary"
	@go clean --mod=mod

.PHONY: setup
## setup: setup go modules
setup:
	@echo "start to download mods"
	@go mod tidy \
		&& @go mod download

.PHONY: env
## env: setup envs
env:
	export GO111MODULES=${GO111MODULES}
	export GOPROXY=${GOPROXY}
	export CGO_ENABLED=${CGO_ENABLED}

.PHONY: test
## test: make test binary
test: clean
	@echo "start to build test app"
	go build --mod=mod -a -o app app.go

.PHONY: help
## help: show help usages
help:
	@echo "this makefile help you build blog binary"
	@echo "Usage: \n"
	@sed -n 's/^##/ /p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
