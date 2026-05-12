package compose

import (
	tea "charm.land/bubbletea/v2"
)

func (m *Model) focussedModel() *tea.Model {
	if m.focus == Primary {
		return m.primary
	}
	return m.secondary
}

func (m *Model) unfocussedModel() *tea.Model {
	if m.focus == Secondary {
		return m.primary
	}
	return m.secondary
}
