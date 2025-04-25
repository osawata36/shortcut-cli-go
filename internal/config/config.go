package config

import (
	"fmt"
	"os"
)

// Config はアプリケーションの設定を表す構造体です
type Config struct {
	APIToken string
}

// Load は環境変数から設定を読み込みます
func Load() (*Config, error) {
	apiToken := os.Getenv("SHORTCUT_API_TOKEN")
	if apiToken == "" {
		return nil, fmt.Errorf("SHORTCUT_API_TOKEN environment variable is not set")
	}

	return &Config{
		APIToken: apiToken,
	}, nil
}
