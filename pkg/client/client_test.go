package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

// mockHTTPClient はテスト用のHTTPクライアントモックです
type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

// mockResponse はテスト用のレスポンスを生成します
func mockResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func TestGetEpic(t *testing.T) {
	tests := []struct {
		name       string
		epicID     int
		response   string
		statusCode int
		wantErr    bool
	}{
		{
			name:   "正常系",
			epicID: 123,
			response: `{
				"id": 123,
				"name": "Test Epic",
				"description": "Test Description",
				"state": "in progress",
				"created_at": "2024-01-01T00:00:00Z",
				"updated_at": "2024-01-01T00:00:00Z"
			}`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "エラー系: 404",
			epicID:     999,
			response:   `{"error": "Not found"}`,
			statusCode: http.StatusNotFound,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockHTTPClient{
				doFunc: func(req *http.Request) (*http.Response, error) {
					return mockResponse(tt.statusCode, tt.response), nil
				},
			}

			client := &clientImpl{
				httpClient: mock,
				apiToken:   "test-token",
			}

			epic, err := client.GetEpic(context.Background(), tt.epicID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && epic == nil {
				t.Error("GetEpic() returned nil epic")
			}
		})
	}
}

func TestSearchStories(t *testing.T) {
	tests := []struct {
		name       string
		params     *SearchStoryParams
		response   string
		statusCode int
		wantErr    bool
	}{
		{
			name: "正常系",
			params: &SearchStoryParams{
				Query: "test",
			},
			response: `[{
				"id": 123,
				"name": "Test Story",
				"description": "Test Description",
				"story_type": "feature",
				"workflow_state_id": "in progress",
				"created_at": "2024-01-01T00:00:00Z",
				"updated_at": "2024-01-01T00:00:00Z"
			}]`,
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name: "エラー系: 429",
			params: &SearchStoryParams{
				Query: "test",
			},
			response:   `{"error": "Too many requests"}`,
			statusCode: http.StatusTooManyRequests,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockHTTPClient{
				doFunc: func(req *http.Request) (*http.Response, error) {
					return mockResponse(tt.statusCode, tt.response), nil
				},
			}

			client := &clientImpl{
				httpClient: mock,
				apiToken:   "test-token",
			}

			stories, err := client.SearchStories(context.Background(), tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchStories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(stories) == 0 {
				t.Error("SearchStories() returned empty slice")
			}
		})
	}
}
