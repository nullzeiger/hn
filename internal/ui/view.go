// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Render
	labelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D7D7D")).Render
	detailStyle = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder())
)

func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("\n  Error: %v\n\n  Press q to exit.\n", m.Err)
	}

	if m.Loading {
		return fmt.Sprintf("\n %s Loading news from Hacker News...\n", m.Spinner.View())
	}

	leftWidth := int(float64(m.Width) * 0.4)
	if leftWidth < 20 {
		leftWidth = m.Width / 2
	}
	rightWidth := m.Width - leftWidth - 6

	left := m.L.View()

	var right string
	selectedItem := m.L.SelectedItem()

	if selectedItem != nil {
		if i, ok := selectedItem.(ListItem); ok {
			s := i.Story
			title := titleStyle(s.Title)
			ts := time.Unix(s.Time, 0).Format("2006-01-02 15:04")

			url := s.URL
			if url == "" {
				url = "<no url>"
			}

			rightBody := fmt.Sprintf("%s\n\n%s %d / %d\n\n%s %s\n\n%s %d punti\n\n%s %s\n\n%s %s",
				title,
				labelStyle("Index:"), i.Index+1, len(m.Stories),
				labelStyle("Author:"), s.By,
				labelStyle("Score:"), s.Score,
				labelStyle("Time:"), ts,
				labelStyle("Link:"), url,
			)
			right = detailStyle.Width(rightWidth).Render(rightBody)
		}
	} else {
		right = detailStyle.Width(rightWidth).Render("No results found..")
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, left, "    ", right)
	header := lipgloss.NewStyle().Bold(true).Render("Hacker News TUI - Press q to exit.")

	return header + "\n\n" + row
}
