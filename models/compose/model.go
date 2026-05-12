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

// --- Model ---

type Model struct {
	primary, secondary         *tea.Model
	primarySize, secondarySize util.Size

	Orientation  Orientation
	ReverseOrder bool
	GetSizes     func(primary, scondary *tea.Model, focus Focus) (util.Size, util.Size)

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
		cmd = (*m.primary).Init()
	}
	if m.secondary != nil {
		cmd = tea.Batch(cmd, (*m.secondary).Init())
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

		if m.focus != msg.focus {
			// change focus
			var cmd tea.Cmd
			if focussedModel := m.focussedModel(); focussedModel != nil {
				util.TryBlurTeaModel(*focussedModel)
			}
			if unfocussedModel := m.unfocussedModel(); unfocussedModel != nil {
				util.TryFocusTeaModel(*unfocussedModel, m.lastKeyMap)
			}
			m.focus = msg.focus
			// and refrsh sizes
			return m, tea.Sequence(cmd, m.refreshSizes())
		}
	}

	// pass remaining messages to focussed model
	if focussedModel := m.focussedModel(); focussedModel != nil {
		return m, util.UpdateTeaModelInplace(msg, focussedModel)
	}

	return m, nil
}

func (m Model) View() tea.View {
	var view1, view2 string
	if m.primary != nil {
		view1 = lipgloss.NewStyle().
			Width(m.primarySize.Width).
			Height(m.primarySize.Height).
			MaxWidth(m.primarySize.Width).
			MaxHeight(m.primarySize.Height).
			Render((*m.primary).View().Content)
	}
	if m.secondary != nil {
		view2 = lipgloss.NewStyle().
			Width(m.secondarySize.Width).
			Height(m.secondarySize.Height).
			MaxWidth(m.secondarySize.Width).
			MaxHeight(m.secondarySize.Height).
			Render((*m.secondary).View().Content)
	}

	if m.ReverseOrder {
		view1, view2 = view2, view1
	}

	var view string
	if m.Orientation == Vertical {
		view = lipgloss.JoinVertical(lipgloss.Left, view1, view2)
	} else {
		view = lipgloss.JoinHorizontal(lipgloss.Top, view1, view2)
	}

	return tea.NewView(view)
}

// --- [util.Focusable] ---

func (m *Model) Focus(parentKeyMap help.KeyMap) tea.Cmd {
	m.lastKeyMap = parentKeyMap

	if focussedModel := m.focussedModel(); focussedModel != nil {
		cmd, _ := util.TryFocusTeaModel(*focussedModel, parentKeyMap)
		return cmd
	}

	return util.AnnounceKeyMapCmd(parentKeyMap)
}

func (m *Model) Blur() {
	if focussedModel := m.focussedModel(); focussedModel != nil {
		util.TryBlurTeaModel(*focussedModel)
	}
}
