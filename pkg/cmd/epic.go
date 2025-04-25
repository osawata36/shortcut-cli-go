package cmd

import (
	"context"
	"fmt"
	"strconv"

	shortcut "github.com/osawata36/shortcut-cli-go/pkg/client"
	"github.com/spf13/cobra"
)

// newEpicCmd はepicコマンドを作成します
func newEpicCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "epic",
		Short: "Epic関連のコマンド",
		Long:  "Epic関連の操作を行うコマンドです。",
	}

	cmd.AddCommand(
		newEpicGetCmd(),
		newEpicListCmd(),
		newEpicSearchCmd(),
	)

	return cmd
}

// newEpicGetCmd はepic getコマンドを作成します
func newEpicGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [epic-id]",
		Short: "Epicの情報を取得",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			epicID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid epic ID: %s", args[0])
			}

			epic, err := client.GetEpic(context.Background(), epicID)
			if err != nil {
				return err
			}

			// 結果の表示
			fmt.Printf("ID: %d\n", epic.ID)
			fmt.Printf("Name: %s\n", epic.Name)
			fmt.Printf("Description: %s\n", epic.Description)
			fmt.Printf("State: %s\n", epic.State)
			fmt.Printf("Created: %s\n", epic.CreatedAt)
			fmt.Printf("Updated: %s\n", epic.UpdatedAt)

			return nil
		},
	}
}

// newEpicListCmd はepic listコマンドを作成します
func newEpicListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Epic一覧を取得",
		RunE: func(cmd *cobra.Command, args []string) error {
			epics, err := client.ListEpics(context.Background())
			if err != nil {
				return err
			}

			// 結果の表示
			for _, epic := range epics {
				fmt.Printf("ID: %d, Name: %s, State: %s\n", epic.ID, epic.Name, epic.State)
			}

			return nil
		},
	}
}

// newEpicSearchCmd はepic searchコマンドを作成します
func newEpicSearchCmd() *cobra.Command {
	var (
		state     string
		ownerID   string
		createdAt string
		updatedAt string
	)

	cmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Epicを検索",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			query := ""
			if len(args) > 0 {
				query = args[0]
			}

			params := &shortcut.SearchEpicParams{
				Query:     query,
				State:     state,
				OwnerID:   ownerID,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}

			epics, err := client.SearchEpics(context.Background(), params)
			if err != nil {
				return err
			}

			// 結果の表示
			for _, epic := range epics {
				fmt.Printf("ID: %d, Name: %s, State: %s\n", epic.ID, epic.Name, epic.State)
			}

			return nil
		},
	}

	// フラグの設定
	cmd.Flags().StringVar(&state, "state", "", "Epicの状態でフィルタ (unstarted/started/done)")
	cmd.Flags().StringVar(&ownerID, "owner", "", "オーナーIDでフィルタ")
	cmd.Flags().StringVar(&createdAt, "created", "", "作成日でフィルタ (YYYY-MM-DD)")
	cmd.Flags().StringVar(&updatedAt, "updated", "", "更新日でフィルタ (YYYY-MM-DD)")

	return cmd
}
