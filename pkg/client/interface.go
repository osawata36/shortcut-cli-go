package client

import "context"

// Client はShortcut APIクライアントのインターフェースです
type Client interface {
	// Epic関連
	GetEpic(ctx context.Context, epicID int) (*Epic, error)
	ListEpics(ctx context.Context) ([]*Epic, error)
	SearchEpics(ctx context.Context, query *SearchEpicParams) ([]*Epic, error)

	// Story関連
	GetStory(ctx context.Context, storyID int) (*Story, error)
	SearchStories(ctx context.Context, query *SearchStoryParams) ([]*Story, error)
}

// Epic はEpicの情報を表す構造体です
type Epic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Story はStoryの情報を表す構造体です
type Story struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	StoryType     string `json:"story_type"`
	WorkflowState State  `json:"workflow_state"`
	EpicID        *int   `json:"epic_id,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// State はワークフローの状態を表す構造体です
type State struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SearchEpicParams はEpic検索のパラメータを表す構造体です
type SearchEpicParams struct {
	Query     string `json:"query,omitempty"`
	State     string `json:"state,omitempty"`
	OwnerID   string `json:"owner,omitempty"`
	CreatedAt string `json:"created,omitempty"`
	UpdatedAt string `json:"updated,omitempty"`
}

// SearchStoryParams はStory検索のパラメータを表す構造体です
type SearchStoryParams struct {
	Query     string `json:"query,omitempty"`
	EpicID    int    `json:"epic,omitempty"`
	State     string `json:"state,omitempty"`
	OwnerID   string `json:"owner,omitempty"`
	StoryType string `json:"type,omitempty"`
	CreatedAt string `json:"created,omitempty"`
	UpdatedAt string `json:"updated,omitempty"`
}
