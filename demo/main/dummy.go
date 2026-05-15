package main

import (
	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-tea/util"
)

// --- Model ---

type DummyModel struct {
	KeyMap help.KeyMap
}

// [DummyModel] implements [tea.Model]
// *[DummyModel] implements [util.Focusable]
var _ tea.Model = DummyModel{}
var _ util.Focusable = (*DummyModel)(nil)

// --- [tea.Model] ---

func (m DummyModel) Init() tea.Cmd { return nil }

func (m DummyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }

func (m DummyModel) View() tea.View {
	return tea.NewView("Dummy View")
}

// --- [util.Focusable] ---

func (m *DummyModel) Focus(parentKeyMap help.KeyMap) tea.Cmd {
	return util.AnnounceKeyMapCmd(parentKeyMap, m.KeyMap)
}

func (m *DummyModel) Blur() {}
