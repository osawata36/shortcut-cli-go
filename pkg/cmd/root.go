package cmd

import (
	"fmt"

	"github.com/osawata36/shortcut-cli-go/internal/config"
	shortcut "github.com/osawata36/shortcut-cli-go/pkg/client"
	"github.com/spf13/cobra"
)

var (
	cfg    *config.Config
	client shortcut.Client

	// バージョン情報
	version   string
	buildTime string
)

// rootCmd はルートコマンドを表します
var rootCmd = &cobra.Command{
	Use:   "shortcut-cli",
	Short: "Shortcut CLI tool",
	Long: `Shortcut CLI tool は、ShortcutのAPIをコマンドラインから操作するためのツールです。
Epic情報の参照やStoryの検索などが行えます。`,
}

// SetVersionInfo はバージョン情報を設定します
func SetVersionInfo(v, bt string) {
	version = v
	buildTime = bt
}

// Execute はルートコマンドを実行します
func Execute(c *config.Config) error {
	cfg = c
	client = shortcut.NewClient(cfg.APIToken)

	// バージョンコマンドの追加
	rootCmd.AddCommand(
		newVersionCmd(),
		newEpicCmd(),
		newStoryCmd(),
	)

	return rootCmd.Execute()
}

// newVersionCmd はバージョン情報を表示するコマンドを作成します
func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "バージョン情報を表示",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\nBuild Time: %s\n", version, buildTime)
		},
	}
}
