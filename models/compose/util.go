package compose

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/jannes-dailidow/go-tea/util"
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

func (m *Model) refreshSizes() tea.Cmd {
	if m.GetSizes != nil {
		m.primarySize, m.secondarySize = m.GetSizes(m.primary, m.secondary, m.focus)
	} else {
		// render view of focussed model
		var focussedView string
		if focussedModel := m.focussedModel(); focussedModel != nil {
			focussedView = (*focussedModel).View().Content
		}

		// calculate sizes
		var focussedSize, unfocussedSize util.Size
		if m.Orientation == Vertical {
			height := min(lipgloss.Height(focussedView), m.size.Height)
			focussedSize = util.Size{m.size.Width, height}
			unfocussedSize = util.Size{m.size.Width, m.size.Height - height}
		} else {
			width := min(lipgloss.Width(focussedView), m.size.Width)
			focussedSize = util.Size{width, m.size.Height}
			unfocussedSize = util.Size{m.size.Width - width, m.size.Height}
		}

		// cache sizes
		if m.focus == Primary {
			m.primarySize, m.secondarySize = focussedSize, unfocussedSize
		} else {
			m.primarySize, m.secondarySize = unfocussedSize, focussedSize
		}
	}

	// apply sizes to models
	var primaryCmd, secondaryCmd tea.Cmd
	if m.primary != nil {
		primaryCmd = util.UpdateTeaModelInplace(m.primarySize, m.primary)
	}
	if m.secondary != nil {
		secondaryCmd = util.UpdateTeaModelInplace(m.secondarySize, m.secondary)
	}

	return tea.Batch(primaryCmd, secondaryCmd)
}
