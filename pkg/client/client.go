package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

const (
	baseURL        = "https://api.app.shortcut.com/api/v3"
	apiTokenHeader = "Shortcut-Token"
)

// httpClient はHTTPクライアントのインターフェースです
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// clientImpl はClient interfaceの実装です
type clientImpl struct {
	httpClient  httpClient
	apiToken    string
	rateLimiter *rate.Limiter
}

// NewClient は新しいShortcutクライアントを作成します
func NewClient(apiToken string) Client {
	// レートリミッターを作成（200リクエスト/分）
	limiter := rate.NewLimiter(rate.Every(time.Minute/200), 1)

	return &clientImpl{
		httpClient:  &http.Client{Timeout: 30 * time.Second},
		apiToken:    apiToken,
		rateLimiter: limiter,
	}
}

// doRequest はHTTPリクエストを実行し、レスポンスを返します
func (c *clientImpl) doRequest(ctx context.Context, method, path string, query url.Values) ([]byte, error) {
	// レートリミッターの待機
	err := c.rateLimiter.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("rate limiter wait: %w", err)
	}

	// URLの構築
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse url: %w", err)
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	// リクエストの作成
	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// ヘッダーの設定
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(apiTokenHeader, c.apiToken)

	// リクエストの実行
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	// レスポンスの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	// ステータスコードの確認
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusTooManyRequests {
			// レートリミットに達した場合は少し待ってリトライ
			time.Sleep(time.Second)
			return c.doRequest(ctx, method, path, query)
		}
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// GetEpic はEpicの情報を取得します
func (c *clientImpl) GetEpic(ctx context.Context, epicID int) (*Epic, error) {
	body, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/epics/%d", epicID), nil)
	if err != nil {
		return nil, err
	}

	var epic Epic
	if err := json.Unmarshal(body, &epic); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &epic, nil
}

// ListEpics は全てのEpicのリストを取得します
func (c *clientImpl) ListEpics(ctx context.Context) ([]*Epic, error) {
	body, err := c.doRequest(ctx, http.MethodGet, "/epics", nil)
	if err != nil {
		return nil, err
	}

	var epics []*Epic
	if err := json.Unmarshal(body, &epics); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return epics, nil
}

// SearchEpics は条件に合うEpicを検索します
func (c *clientImpl) SearchEpics(ctx context.Context, params *SearchEpicParams) ([]*Epic, error) {
	query := url.Values{}
	if params.Query != "" {
		query.Set("query", params.Query)
	}
	if params.State != "" {
		query.Set("state", params.State)
	}
	if params.OwnerID != "" {
		query.Set("owner", params.OwnerID)
	}
	if params.CreatedAt != "" {
		query.Set("created", params.CreatedAt)
	}
	if params.UpdatedAt != "" {
		query.Set("updated", params.UpdatedAt)
	}

	body, err := c.doRequest(ctx, http.MethodGet, "/search/epics", query)
	if err != nil {
		return nil, err
	}

	var epics []*Epic
	if err := json.Unmarshal(body, &epics); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return epics, nil
}

// GetStory はStoryの情報を取得します
func (c *clientImpl) GetStory(ctx context.Context, storyID int) (*Story, error) {
	body, err := c.doRequest(ctx, http.MethodGet, fmt.Sprintf("/stories/%d", storyID), nil)
	if err != nil {
		return nil, err
	}

	var story Story
	if err := json.Unmarshal(body, &story); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return &story, nil
}

// SearchStories は条件に合うStoryを検索します
func (c *clientImpl) SearchStories(ctx context.Context, params *SearchStoryParams) ([]*Story, error) {
	query := url.Values{}
	if params.Query != "" {
		query.Set("query", params.Query)
	}
	if params.EpicID != 0 {
		query.Set("epic", fmt.Sprintf("%d", params.EpicID))
	}
	if params.State != "" {
		query.Set("state", params.State)
	}
	if params.OwnerID != "" {
		query.Set("owner", params.OwnerID)
	}
	if params.StoryType != "" {
		query.Set("type", params.StoryType)
	}
	if params.CreatedAt != "" {
		query.Set("created", params.CreatedAt)
	}
	if params.UpdatedAt != "" {
		query.Set("updated", params.UpdatedAt)
	}

	body, err := c.doRequest(ctx, http.MethodGet, "/search/stories", query)
	if err != nil {
		return nil, err
	}

	var stories []*Story
	if err := json.Unmarshal(body, &stories); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return stories, nil
}
