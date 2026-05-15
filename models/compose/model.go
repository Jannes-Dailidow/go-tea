package compose

import (
	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

const (
	Primary   Focus = false
	Secondary Focus = true

	Vertical   Orientation = false
	Horizontal Orientation = true
)

type (
	Focus       bool
	Orientation bool
)

type SizeResolver func(size util.Size, primary, secondary tea.Model, focus Focus, orientation Orientation) (util.Size, util.Size)
type MsgRoute func(msg tea.Msg) (bool, Focus)

// --- Model ---

type Model struct {
	primary, secondary         tea.Model
	primarySize, secondarySize util.Size

	Orientation  Orientation
	ReverseOrder bool
	SizeResolver SizeResolver
	MsgRoutes    []MsgRoute

	id         util.ModelId
	size       util.Size
	focussed   bool
	focus      Focus
	lastKeyMap help.KeyMap
}

// [Model] implements [tea.Model]
// *[Model] implements [util.Focusable]
var _ tea.Model = Model{}
var _ util.Focusable = (*Model)(nil)

// --- [tea.Model] ---

func (m Model) Init() tea.Cmd {
	var cmd tea.Cmd
	if m.primary != nil {
		cmd = m.primary.Init()
	}
	if m.secondary != nil {
		cmd = tea.Batch(cmd, m.secondary.Init())
	}
	if m.focussed {
		if focussedModel := m.focussedModel(); focussedModel != nil {
			focusCmd, _ := util.TryFocusTeaModel(*focussedModel, nil)
			cmd = tea.Sequence(cmd, focusCmd)
		}
	}
	return cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.size.UpdateFromMsg(msg) {
		return m, m.refreshSizes()
	}

	switch msg := msg.(type) {
	case ChangeFocusMsg:
		if msg.id != m.id {
			break
		}

		return m, m.ChangeFocus(msg.focus)
	}

	// handle custom MsgRoutes
	for _, msgRoute := range m.MsgRoutes {
		if ok, focus := msgRoute(msg); ok {
			model := m.modelByFocus(focus)
			if model == nil {
				return m, nil
			}
			return m, util.TeaUpdateModelInplace(msg, model)
		}
	}

	// pass remaining messages to focussed model
	if model := m.focussedModel(); model != nil {
		return m, util.TeaUpdateModelInplace(msg, model)
	}

	return m, nil
}

func (m Model) View() tea.View {
	primarySize, secondarySize := m.resolveSizes()

	var primaryView, secondaryView tea.View
	if m.primary != nil {
		primaryView = m.primary.View()
		primaryView.Content = lipgloss.NewStyle().
			Width(primarySize.Width).
			Height(primarySize.Height).
			MaxWidth(primarySize.Width).
			MaxHeight(primarySize.Height).
			Render(primaryView.Content)
	}
	if m.secondary != nil {
		secondaryView = m.secondary.View()
		secondaryView.Content = lipgloss.NewStyle().
			Width(secondarySize.Width).
			Height(secondarySize.Height).
			MaxWidth(secondarySize.Width).
			MaxHeight(secondarySize.Height).
			Render(secondaryView.Content)
	}

	// merge views based on focus
	var view tea.View
	if m.focus == Primary {
		view = util.TeaMergeViews(primaryView, secondaryView)
	} else {
		view = util.TeaMergeViews(secondaryView, primaryView)
	}

	// swap views when ReverseOrder is set
	if m.ReverseOrder {
		primaryView, secondaryView = secondaryView, primaryView
	}

	// join views depending on Orientation
	if m.Orientation == Vertical {
		view.SetContent(lipgloss.JoinVertical(lipgloss.Left, primaryView.Content, secondaryView.Content))
	} else {
		view.SetContent(lipgloss.JoinHorizontal(lipgloss.Top, primaryView.Content, secondaryView.Content))
	}

	return view
}

// --- [util.Focusable] ---

func (m *Model) Focus(parentKeyMap help.KeyMap) tea.Cmd {
	m.lastKeyMap = parentKeyMap

	if model := m.focussedModel(); model != nil {
		cmd, ok := util.TryFocusTeaModel(*model, parentKeyMap)
		if ok {
			return cmd
		}
	}

	// fallback to announce keymap yourself
	return util.AnnounceKeyMapCmd(parentKeyMap)
}

func (m *Model) Blur() {
	if model := m.focussedModel(); model != nil {
		_ = util.TryBlurTeaModel(*model)
	}
}

func (m *Model) ChangeFocus(focus Focus) tea.Cmd {
	if m.focus == focus {
		return nil
	}

	m.focus = focus

	if !m.focussed {
		return m.refreshSizes()
	}

	if model := m.unfocussedModel(); model != nil {
		_ = util.TryBlurTeaModel(*model)
	}
	if model := m.focussedModel(); model != nil {
		cmd, ok := util.TryFocusTeaModel(*model, m.lastKeyMap)
		if ok {
			return tea.Sequence(cmd, m.refreshSizes())
		}
	}

	// fallback to announce keymap yourself
	return tea.Sequence(util.AnnounceKeyMapCmd(m.lastKeyMap), m.refreshSizes())
}
