package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config はアプリケーションの設定を表す構造体です
type Config struct {
	APIToken string
}

// Load は.envファイルと環境変数から設定を読み込みます
func Load() (*Config, error) {
	// .envファイルの読み込み（ファイルが存在しない場合はスキップ）
	_ = godotenv.Load()

	apiToken := os.Getenv("SHORTCUT_API_TOKEN")
	if apiToken == "" {
		return nil, fmt.Errorf("SHORTCUT_API_TOKEN environment variable is not set")
	}

	return &Config{
		APIToken: apiToken,
	}, nil
}
