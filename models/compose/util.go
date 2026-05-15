package compose

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

func (m *Model) modelByFocus(focus Focus) *tea.Model {
	if focus == Primary {
		return &m.primary
	}
	return &m.secondary
}

func (m *Model) focussedModel() *tea.Model {
	return m.modelByFocus(m.focus)
}

func (m *Model) unfocussedModel() *tea.Model {
	return m.modelByFocus(!m.focus)
}

func (m *Model) resolveSizes() (util.Size, util.Size) {
	if m.SizeResolver != nil {
		return m.SizeResolver(m.size, m.primary, m.secondary, m.focus, m.Orientation)
	} else {
		return defaultSizeResolver(m.size, m.primary, m.secondary, m.focus, m.Orientation)
	}
}

func (m *Model) refreshSizes() tea.Cmd {
	// resolve sizes
	primarySize, secondarySize := m.resolveSizes()

	// apply sizes to models if they changed
	var primaryCmd, secondaryCmd tea.Cmd
	if m.primary != nil && m.primarySize != primarySize {
		m.primarySize = primarySize
		primaryCmd = util.TeaUpdateModelInplace(primarySize.ToMsg(), &m.primary)
	}
	if m.secondary != nil && m.secondarySize != secondarySize {
		m.secondarySize = secondarySize
		secondaryCmd = util.TeaUpdateModelInplace(secondarySize.ToMsg(), &m.secondary)
	}

	return tea.Batch(primaryCmd, secondaryCmd)
}

// [defaultSizeResolver] implements [SizeResolver]
var _ SizeResolver = defaultSizeResolver

func defaultSizeResolver(size util.Size, primary, secondary tea.Model, focus Focus, orientation Orientation) (util.Size, util.Size) {
	// render view of focussed model
	var focussedView string
	if focus == Primary && primary != nil {
		focussedView = primary.View().Content
	}
	if focus == Secondary && secondary != nil {
		focussedView = secondary.View().Content
	}

	// calculate sizes
	var focussedSize, unfocussedSize util.Size
	if orientation == Vertical {
		height := min(lipgloss.Height(focussedView), size.Height)
		focussedSize = util.Size{size.Width, height}
		unfocussedSize = util.Size{size.Width, size.Height - height}
	} else {
		width := min(lipgloss.Width(focussedView), size.Width)
		focussedSize = util.Size{width, size.Height}
		unfocussedSize = util.Size{size.Width - width, size.Height}
	}

	if focus == Primary {
		return focussedSize, unfocussedSize
	} else {
		return unfocussedSize, focussedSize
	}
}
