package cmd

import (
	"context"
	"fmt"
	"strconv"

	shortcut "github.com/osawata36/shortcut-cli-go/pkg/client"
	"github.com/spf13/cobra"
)

// newStoryCmd はstoryコマンドを作成します
func newStoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "story",
		Short: "Story関連のコマンド",
		Long:  "Story関連の操作を行うコマンドです。",
	}

	cmd.AddCommand(
		newStoryGetCmd(),
		newStorySearchCmd(),
	)

	return cmd
}

// newStoryGetCmd はstory getコマンドを作成します
func newStoryGetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [story-id]",
		Short: "Storyの情報を取得",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			storyID, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid story ID: %s", args[0])
			}

			story, err := client.GetStory(context.Background(), storyID)
			if err != nil {
				return err
			}

			// 結果の表示
			fmt.Printf("ID: %d\n", story.ID)
			fmt.Printf("Name: %s\n", story.Name)
			fmt.Printf("Description: %s\n", story.Description)
			fmt.Printf("Type: %s\n", story.StoryType)
			fmt.Printf("State: %s\n", story.State)
			if story.EpicID != nil {
				fmt.Printf("Epic ID: %d\n", *story.EpicID)
			}
			fmt.Printf("Created: %s\n", story.CreatedAt)
			fmt.Printf("Updated: %s\n", story.UpdatedAt)

			return nil
		},
	}
}

// newStorySearchCmd はstory searchコマンドを作成します
func newStorySearchCmd() *cobra.Command {
	var (
		epicID    int
		state     string
		ownerID   string
		storyType string
		createdAt string
		updatedAt string
	)

	cmd := &cobra.Command{
		Use:   "search [query]",
		Short: "Storyを検索",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			query := ""
			if len(args) > 0 {
				query = args[0]
			}

			params := &shortcut.SearchStoryParams{
				Query:     query,
				EpicID:    epicID,
				State:     state,
				OwnerID:   ownerID,
				StoryType: storyType,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}

			stories, err := client.SearchStories(context.Background(), params)
			if err != nil {
				return err
			}

			// 結果の表示
			for _, story := range stories {
				fmt.Printf("ID: %d, Name: %s, Type: %s, State: %s\n",
					story.ID, story.Name, story.StoryType, story.State)
			}

			return nil
		},
	}

	// フラグの設定
	cmd.Flags().IntVar(&epicID, "epic", 0, "Epic IDでフィルタ")
	cmd.Flags().StringVar(&state, "state", "", "Storyの状態でフィルタ")
	cmd.Flags().StringVar(&ownerID, "owner", "", "オーナーIDでフィルタ")
	cmd.Flags().StringVar(&storyType, "type", "", "Storyタイプでフィルタ (feature/bug/chore)")
	cmd.Flags().StringVar(&createdAt, "created", "", "作成日でフィルタ (YYYY-MM-DD)")
	cmd.Flags().StringVar(&updatedAt, "updated", "", "更新日でフィルタ (YYYY-MM-DD)")

	return cmd
}
