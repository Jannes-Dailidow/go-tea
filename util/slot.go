package util

import tea "charm.land/bubbletea/v2"

type Slot[M tea.Model] struct {
	Model M
}

func (s Slot[M]) Init() tea.Cmd { return s.Model.Init() }

func (s Slot[M]) Update(tea.Msg) (tea.Model, tea.Cmd)

func (s Slot[M]) View() tea.View { return s.Model.View() }
