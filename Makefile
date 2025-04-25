.PHONY: build test clean lint

# バイナリ名
BINARY_NAME=shortcut-cli

# Go関連のコマンド
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# ビルド対象のメインパッケージ
MAIN_PACKAGE=cmd/shortcut/main.go

# 環境変数
VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# デフォルトターゲット
all: lint test build

# ビルド
build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)

# テスト実行
test:
	$(GOTEST) -v ./...

# 依存関係の更新
deps:
	$(GOGET) -u
	$(GOMOD) tidy

# クリーンアップ
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME).exe

# リント
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

# インストール
install:
	go install $(LDFLAGS) 