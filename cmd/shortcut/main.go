package main

import (
	"fmt"
	"os"

	"github.com/osawata36/shortcut-cli-go/internal/config"
	"github.com/osawata36/shortcut-cli-go/pkg/cmd"
)

// これらの変数はビルド時に-ldflagsで設定されます
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	// 設定の読み込み
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "設定の読み込みに失敗しました: %v\n", err)
		os.Exit(1)
	}

	// バージョン情報を設定
	cmd.SetVersionInfo(Version, BuildTime)

	// ルートコマンドの実行
	if err := cmd.Execute(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
}
