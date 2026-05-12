package util

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func TeaMsgToCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}

func UpdateTeaModelInplace[M any](msg tea.Msg, model *M) tea.Cmd {
	var modelAny any = *model

	if modelUpdatable, ok := modelAny.(updatableSelf[M]); ok {
		modelUpdated, cmd := modelUpdatable.Update(msg)
		*model = modelUpdated
		return cmd
	}

	if modelUpdatable, ok := modelAny.(updatableTea); ok {
		modelUpdated, cmd := modelUpdatable.Update(msg)
		if modelUpdated, ok := modelUpdated.(M); ok {
			*model = modelUpdated
			return cmd
		}
		return cmd
	}

	panic(fmt.Sprintf("no supported update method in provided type %T", &model))
}

type updatableTea interface {
	Update(tea.Msg) (tea.Model, tea.Cmd)
}
type updatableSelf[T any] interface {
	Update(tea.Msg) (T, tea.Cmd)
}
