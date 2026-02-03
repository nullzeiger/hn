// Copyright 2026 Ivan Guerreschi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle  = lipgloss.NewStyle().Bold(true).Render
	LabelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D7D7D")).Render
	DetailStyle = lipgloss.NewStyle().Padding(1).Border(lipgloss.NormalBorder())
)
