package api

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type mockRoundTripper struct {
	handler func(req *http.Request) *http.Response
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.handler(req), nil
}

func newMockResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
		Header:     make(http.Header),
	}
}

func TestFetchStories_Success(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = &http.Client{
		Transport: &mockRoundTripper{
			handler: func(req *http.Request) *http.Response {
				switch req.URL.String() {
				case TopStoriesURL:
					return newMockResponse(200, `[1,2]`)
				case "https://hacker-news.firebaseio.com/v0/item/1.json":
					return newMockResponse(200, `{
						"id":1,
						"title":"Story One",
						"url":"https://example.com/1",
						"score":100,
						"by":"author1",
						"time":123456
					}`)
				case "https://hacker-news.firebaseio.com/v0/item/2.json":
					return newMockResponse(200, `{
						"id":2,
						"title":"Story Two",
						"url":"https://example.com/2",
						"score":200,
						"by":"author2",
						"time":654321
					}`)
				default:
					return newMockResponse(404, "")
				}
			},
		},
	}

	stories, err := FetchStories()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(stories) != 2 {
		t.Fatalf("expected 2 stories, got %d", len(stories))
	}

	if stories[0].Title != "Story One" {
		t.Errorf("expected 'Story One', got %q", stories[0].Title)
	}

	if stories[1].Title != "Story Two" {
		t.Errorf("expected 'Story Two', got %q", stories[1].Title)
	}
}

func TestFetchStories_ItemErrorFallback(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = &http.Client{
		Transport: &mockRoundTripper{
			handler: func(req *http.Request) *http.Response {
				switch req.URL.String() {
				case TopStoriesURL:
					return newMockResponse(200, `[42]`)
				case "https://hacker-news.firebaseio.com/v0/item/42.json":
					return newMockResponse(200, `invalid-json`)
				default:
					return newMockResponse(404, "")
				}
			},
		},
	}

	stories, err := FetchStories()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(stories) != 1 {
		t.Fatalf("expected 1 story, got %d", len(stories))
	}

	if stories[0].Title != "[error fetching]" {
		t.Errorf("expected fallback title, got %q", stories[0].Title)
	}
}

func TestFetchStories_TopStoriesError(t *testing.T) {
	originalClient := httpClient
	defer func() { httpClient = originalClient }()

	httpClient = &http.Client{
		Transport: &mockRoundTripper{
			handler: func(req *http.Request) *http.Response {
				return newMockResponse(500, "")
			},
		},
	}

	_, err := FetchStories()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
