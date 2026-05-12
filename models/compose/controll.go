package compose

import (
	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

type Controll struct {
	id util.ModelId
}

func (c Controll) ChangeFocus(focus Focus) tea.Msg {
	return ChangeFocusMsg{c.id, focus}
}

func (c Controll) ChangeFocusPrimary() tea.Msg {
	return ChangeFocusMsg{c.id, Primary}
}

func (c Controll) ChangeFocusSecondary() tea.Msg {
	return ChangeFocusMsg{c.id, Secondary}
}
