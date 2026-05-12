package compose

import (
	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

type ChangeFocusMsg struct {
	Id    util.ModelId
	Focus Focus
}

func ChangeFocus(id util.ModelId, focus Focus) tea.Msg {
	return ChangeFocusMsg{id, focus}
}

func ChangeFocusPrimary(id util.ModelId) tea.Msg {
	return ChangeFocusMsg{id, Primary}
}

func ChangeFocusSecondary(id util.ModelId) tea.Msg {
	return ChangeFocusMsg{id, Secondary}
}
