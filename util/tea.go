package util

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-slicest"
)

func TeaMsgToCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}

func TeaUpdateModelInplace[M any](msg tea.Msg, model *M) tea.Cmd {
	switch updatable := any(*model).(type) {
	case updatableSelf[M]:
		modelUpdated, cmd := updatable.Update(msg)
		*model = modelUpdated
		return cmd

	case updatableTea:
		modelUpdated, cmd := updatable.Update(msg)
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
type updatableSelf[M any] interface {
	Update(tea.Msg) (M, tea.Cmd)
}

func TeaMergeViews(views ...tea.View) tea.View {
	switch len(views) {
	case 0:
		return tea.NewView("")
	case 1:
		return views[0]
	default:
		return slicest.ReduceD(views[1:], views[0], func(view, viewResult tea.View) tea.View {
			if viewResult.AltScreen == false {
				viewResult.AltScreen = view.AltScreen
			}
			if viewResult.BackgroundColor == nil {
				viewResult.BackgroundColor = view.BackgroundColor
			}
			if viewResult.Content == "" {
				viewResult.Content = view.Content
			}
			if viewResult.Cursor == nil {
				viewResult.Cursor = view.Cursor
			}
			if viewResult.DisableBracketedPasteMode == false {
				viewResult.DisableBracketedPasteMode = view.DisableBracketedPasteMode
			}
			if viewResult.ForegroundColor == nil {
				viewResult.ForegroundColor = view.ForegroundColor
			}
			if viewResult.ProgressBar == nil {
				viewResult.ProgressBar = view.ProgressBar
			}
			if viewResult.ReportFocus == false {
				viewResult.ReportFocus = view.ReportFocus
			}
			if viewResult.WindowTitle == "" {
				viewResult.WindowTitle = view.WindowTitle
			}

			if viewResult.OnMouse != nil {
				if view.OnMouse != nil {
					viewResult.OnMouse = func(msg tea.MouseMsg) tea.Cmd {
						return tea.Batch(viewResult.OnMouse(msg), view.OnMouse(msg))
					}
				}
				viewResult.OnMouse = view.OnMouse
			}

			viewResult.MouseMode = max(viewResult.MouseMode, view.MouseMode)

			if viewResult.KeyboardEnhancements.ReportEventTypes == false {
				viewResult.KeyboardEnhancements.ReportEventTypes = view.KeyboardEnhancements.ReportEventTypes
			}
			if viewResult.KeyboardEnhancements.ReportAlternateKeys == false {
				viewResult.KeyboardEnhancements.ReportAlternateKeys = view.KeyboardEnhancements.ReportAlternateKeys
			}
			if viewResult.KeyboardEnhancements.ReportAllKeysAsEscapeCodes == false {
				viewResult.KeyboardEnhancements.ReportAllKeysAsEscapeCodes = view.KeyboardEnhancements.ReportAllKeysAsEscapeCodes
			}
			if viewResult.KeyboardEnhancements.ReportAssociatedText == false {
				viewResult.KeyboardEnhancements.ReportAssociatedText = view.KeyboardEnhancements.ReportAssociatedText
			}
			return viewResult
		})
	}
}
