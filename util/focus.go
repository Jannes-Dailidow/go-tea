package util

import (
	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
)

type Focusable interface {
	Focus(help.KeyMap) tea.Cmd
	Blur()
}

func TryFocusTeaModel(m tea.Model, parentKeyMap help.KeyMap) (tea.Cmd, bool) {
	if focusable, ok := m.(Focusable); ok {
		return focusable.Focus(parentKeyMap), true
	} else {
		return AnnounceKeyMapCmd(parentKeyMap), false
	}
}

func TryBlurTeaModel(m tea.Model) bool {
	if focusable, ok := m.(Focusable); ok {
		focusable.Blur()
		return true
	} else {
		return false
	}
}
