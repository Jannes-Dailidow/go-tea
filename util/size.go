package util

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Size struct{ Width, Height int }

func (s *Size) UpdateFromMsg(msg tea.Msg) bool {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		s.Width, s.Height = msg.Width, msg.Height
		return true
	}
	return false
}

func (s Size) ToMsg() tea.WindowSizeMsg {
	return tea.WindowSizeMsg{
		Width:  s.Width,
		Height: s.Height,
	}
}

func SizeFromView(view string) Size {
	return Size{lipgloss.Width(view), lipgloss.Height(view)}
}
