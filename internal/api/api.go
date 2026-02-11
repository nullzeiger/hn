// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	TopStoriesURL = "https://hacker-news.firebaseio.com/v0/topstories.json"
	ItemURL       = "https://hacker-news.firebaseio.com/v0/item/%d.json"
	Limit         = 30
)

type Story struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Score int    `json:"score"`
	By    string `json:"by"`
	Time  int64  `json:"time"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func fetchJSON(url string, target any) error {
	resp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func FetchStories() ([]Story, error) {
	var ids []int
	if err := fetchJSON(TopStoriesURL, &ids); err != nil {
		return nil, err
	}

	count := min(len(ids), Limit)
	ids = ids[:count]

	stories := make([]Story, len(ids))
	var wg sync.WaitGroup

	for i, id := range ids {
		wg.Go(func() {
			var s Story
			url := fmt.Sprintf(ItemURL, id)
			if err := fetchJSON(url, &s); err != nil {
				s = Story{ID: id, Title: "[error fetching]"}
			}
			stories[i] = s
		})
	}

	wg.Wait()
	return stories, nil
}
