// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import (
	"fmt"
	"time"
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
			title := TitleStyle(s.Title)
			ts := time.Unix(s.Time, 0).Format("2006-01-02 15:04")

			url := s.URL
			if url == "" {
				url = "<no url>"
			}

			rightBody := fmt.Sprintf("%s\n\n%s %d / %d\n\n%s %s\n\n%s %d punti\n\n%s %s\n\n%s %s",
				title,
				LabelStyle("Index:"), i.Index+1, len(m.Stories),
				LabelStyle("Author:"), s.By,
				LabelStyle("Score:"), s.Score,
				LabelStyle("Time:"), ts,
				LabelStyle("Link:"), url,
			)
			right = DetailStyle.Width(rightWidth).Render(rightBody)
		}
	} else {
		right = DetailStyle.Width(rightWidth).Render("No results found..")
	}

	row := Row(left, right)
	header := HeaderStyle("Hacker News TUI - Press q to exit.")

	return header + "\n\n" + row
}
