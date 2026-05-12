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
	primary, secondary *tea.Model

	Orientation  Orientation
	ReverseOrder bool

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
		var focussedView string
		if focussedModel := m.focussedModel(); focussedModel != nil {
			focussedView = (*focussedModel).View().Content
		}

		var focussedSize util.Size
		var unfocussedSize util.Size
		if m.Orientation == Vertical {
			height := min(lipgloss.Height(focussedView), m.size.Height)
			focussedSize = util.Size{m.size.Width, height}
			unfocussedSize = util.Size{m.size.Width, m.size.Height - height}
		} else {
			width := min(lipgloss.Width(focussedView), m.size.Width)
			focussedSize = util.Size{width, m.size.Height}
			unfocussedSize = util.Size{m.size.Width - width, m.size.Height}
		}

		var focussedCmd tea.Cmd
		var unfocussedCmd tea.Cmd
		if model := m.focussedModel(); model != nil {
			focussedCmd = util.UpdateTeaModelInplace(focussedSize, model)
		}
		if model := m.unfocussedModel(); model != nil {
			unfocussedCmd = util.UpdateTeaModelInplace(unfocussedSize, model)
		}

		return m, tea.Batch(focussedCmd, unfocussedCmd)
	}

	return m, nil
}

func (m Model) View() tea.View {
	return tea.NewView("")
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
