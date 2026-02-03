// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nullzeiger/hn/internal/api"
)

type ListItem struct {
	Story api.Story
	Index int
}

func (i ListItem) Title() string { return fmt.Sprintf("%2d. %s", i.Index+1, i.Story.Title) }
func (i ListItem) Description() string {
	return fmt.Sprintf("by: %s • %d▲", i.Story.By, i.Story.Score)
}
func (i ListItem) FilterValue() string { return i.Story.Title }

type StoriesLoadedMsg []api.Story
type ErrMsg error

type Model struct {
	L       list.Model
	Stories []api.Story
	Spinner spinner.Model
	Width   int
	Height  int
	Loading bool
	Err     error
}

func InitialModel() Model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.FilterInput.Prompt = " Search: "
	l.FilterInput.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	l.SetShowTitle(false)

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return Model{
		L:       l,
		Spinner: s,
		Loading: true,
	}
}

func FetchStoriesCmd() tea.Msg {
	stories, err := api.FetchStories()
	if err != nil {
		return ErrMsg(err)
	}
	return StoriesLoadedMsg(stories)
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Spinner.Tick, FetchStoriesCmd)
}
