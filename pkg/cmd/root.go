package cmd

import (
	"github.com/osawata36/shortcut-cli-go/internal/config"
	shortcut "github.com/osawata36/shortcut-cli-go/pkg/client"
	"github.com/spf13/cobra"
)

var (
	cfg    *config.Config
	client shortcut.Client
)

// rootCmd はルートコマンドを表します
var rootCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Shortcut CLI tool",
	Long: `Shortcut CLI tool は、ShortcutのAPIをコマンドラインから操作するためのツールです。
Epic情報の参照やStoryの検索などが行えます。`,
}

// Execute はルートコマンドを実行します
func Execute(c *config.Config) error {
	cfg = c
	client = shortcut.NewClient(cfg.APIToken)

	// サブコマンドの追加
	rootCmd.AddCommand(
		newEpicCmd(),
		newStoryCmd(),
	)

	return rootCmd.Execute()
}
