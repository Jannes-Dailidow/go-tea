package keyhelp

import (
	"slices"
	"strings"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m *Model) shortView() tea.View {
	if len(m.KeyMap.ShortHelp()) == 0 {
		return tea.NewView("")
	}

	var b strings.Builder
	var usedWidth int
	var items []string
	separator := m.Styles.ShortSeparator.Inline(true).Render(ShortSeparator)
	tail := " " + m.Styles.Ellipsis.Inline(true).Render(Ellipsis)
	tailLen := lipgloss.Width(tail)

	for i, kb := range m.KeyMap.ShortHelp() {
		if !kb.Enabled() {
			continue
		}

		// Sep
		var sep string
		if i > 0 {
			sep = separator
		}

		// Item
		str := sep +
			m.Styles.ShortKey.Inline(true).Render(kb.Help().Key) + " " +
			m.Styles.ShortDesc.Inline(true).Render(kb.Help().Desc)

		items = append(items, str)
	}

	for i, item := range items {
		itemLen := lipgloss.Width(item)
		if i < len(items)-1 {
			// when not last
			if usedWidth+itemLen+tailLen <= m.size.Width {
				// when next items and at least the tail fit
				usedWidth += itemLen
				b.WriteString(item)
			} else {
				// else just add the tail
				usedWidth += tailLen
				b.WriteString(tail)
				break
			}
		} else {
			// when last
			if usedWidth+itemLen <= m.size.Width {
				// last item fits
				b.WriteString(item)
			} else if usedWidth+tailLen <= m.size.Width {
				// tail fits
				b.WriteString(tail)
			}
			// nothing fits
		}
	}

	return tea.NewView(b.String())
}

func (m *Model) fullView() tea.View {
	if len(m.KeyMap.FullHelp()) == 0 {
		return tea.NewView("")
	}

	var cols []string
	var result []string
	var usedWidth int
	separator := m.Styles.FullSeparator.Inline(true).Render(FullSeparator)
	tail := " " + m.Styles.Ellipsis.Inline(true).Render(Ellipsis)
	tailLen := lipgloss.Width(tail)

	// Iterate over groups to build columns
	for i, group := range m.KeyMap.FullHelp() {
		if group == nil || !slices.ContainsFunc(group, func(binding key.Binding) bool {
			return binding.Enabled()
		}) {
			// ignore groups with unly disabled bindings
			continue
		}
		var (
			sep          string
			keys         []string
			descriptions []string
		)

		// Sep
		if i > 0 {
			sep = separator
		}

		// Separate keys and descriptions into different slices
		for _, binding := range group {
			if !binding.Enabled() {
				// ignore disabled bindings
				continue
			}
			keys = append(keys, binding.Help().Key)
			descriptions = append(descriptions, binding.Help().Desc)
		}

		// Column
		col := lipgloss.JoinHorizontal(lipgloss.Top,
			sep,
			m.Styles.FullKey.Render(lipgloss.JoinVertical(lipgloss.Left, keys...)),
			" ",
			m.Styles.FullDesc.Render(lipgloss.JoinVertical(lipgloss.Left, descriptions...)),
		)

		cols = append(cols, col)
	}

	for i, col := range cols {
		colLen := lipgloss.Width(col)
		if i < len(cols)-1 {
			// when not last
			if usedWidth+colLen+tailLen <= m.size.Width {
				// when next items and at least the tail fit
				usedWidth += colLen
				result = append(result, col)
			} else {
				// else just add the tail
				usedWidth += tailLen
				result = append(result, tail)
				break
			}
		} else {
			// when last
			if usedWidth+colLen <= m.size.Width {
				// last item fits
				result = append(result, col)
			} else if usedWidth+tailLen <= m.size.Width {
				// tail fits
				result = append(result, tail)
			}
			// nothing fits
		}
	}

	return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, result...))
}
