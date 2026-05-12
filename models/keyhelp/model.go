package keyhelp

import (
	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

const (
	ShortSeparator string = " • "
	FullSeparator  string = "    "
	Ellipsis       string = "…"
)

type Model struct {
	KeyMap   help.KeyMap
	Styles   help.Styles
	Expanded bool
	size     util.Size
}

// [Model] implements [tea.Model]
// *[Model] implements [util.Focusable]
var _ tea.Model = Model{}
var _ util.Focusable = (*Model)(nil)

// --- New ---

func New() *Model {
	return &Model{
		Styles: help.DefaultDarkStyles(),
	}
}

// --- [tea.Model] ---

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.size.UpdateFromMsg(msg) {
		return m, nil
	}

	if msg, ok := msg.(util.AnnounceKeyMapMsg); ok {
		m.KeyMap = msg.KeyMap
		return m, nil
	}

	return m, nil
}

func (m Model) View() tea.View {
	if m.KeyMap == nil {
		return tea.NewView("")
	} else if m.Expanded {
		return m.fullView()
	} else {
		return m.shortView()
	}

}

// --- [util.Focusable] ---

func (m *Model) Focus(parentKeyMap help.KeyMap) tea.Cmd {
	return util.AnnounceKeyMapCmd(parentKeyMap)
}

func (m *Model) Blur() {}

// --- other ---

func (m *Model) ToggleExpanded() { m.Expanded = !m.Expanded }
