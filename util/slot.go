package util

import (
	"charm.land/bubbles/v2/help"
	tea "charm.land/bubbletea/v2"
)

type Slot struct {
	Model tea.Model

	Size   Size
	Offset Vec2[int]
}

func (s *Slot) Init() tea.Cmd { return s.Model.Init() }

func (s *Slot) Update(msg tea.Msg) tea.Cmd {
	switch _msg := msg.(type) {
	case tea.MouseMsg:
		if s.isOutbound(_msg.Mouse().X, _msg.Mouse().Y) {
			return nil
		}
		// Spoof mouse position
		msg = s.adaptMouseMsg(_msg)
	case tea.CursorPositionMsg:
		// Spoof cursor position
		_msg.X += s.Offset.X
		_msg.Y += s.Offset.Y
		msg = _msg
	case tea.WindowSizeMsg:
		// Spoof size
		msg = s.Size.ToMsg()
	}

	var cmd tea.Cmd
	s.Model, cmd = s.Model.Update(msg)
	return cmd
}

func (s *Slot) View() tea.View {
	view := s.Model.View()

	// Adjust Cursor position
	if view.Cursor != nil {
		view.Cursor.Position.X += s.Offset.X
		view.Cursor.Position.Y += s.Offset.Y
	}

	// Intercept OnMouse
	if innerOnMouse := view.OnMouse; innerOnMouse != nil {
		view.OnMouse = func(msg tea.MouseMsg) tea.Cmd {
			if s.isOutbound(msg.Mouse().X, msg.Mouse().Y) {
				return nil
			}
			// Spoof mouse position
			return innerOnMouse(s.adaptMouseMsg(msg))
		}
	}

	return view
}

func (s *Slot) isOutbound(x, y int) bool {
	return x < s.Offset.X || y < s.Offset.Y || x >= s.Offset.X+s.Size.Width || y >= s.Offset.Y+s.Size.Height
}

func (s *Slot) adaptMouseMsg(msg tea.MouseMsg) tea.MouseMsg {
	// Spoof mouse position
	mouse := msg.Mouse()
	mouse.X -= s.Offset.X
	mouse.Y -= s.Offset.Y

	// Convert back to original type
	switch msg.(type) {
	case tea.MouseClickMsg:
		return tea.MouseClickMsg(mouse)
	case tea.MouseMotionMsg:
		return tea.MouseMotionMsg(mouse)
	case tea.MouseReleaseMsg:
		return tea.MouseReleaseMsg(mouse)
	case tea.MouseWheelMsg:
		return tea.MouseWheelMsg(mouse)
	default:
		return tea.MouseMotionMsg(mouse)
	}
}

func (s *Slot) TryFocusWithFallback(parentKeyMap help.KeyMap) tea.Cmd {
	if model, ok := s.Model.(Focusable); ok {
		return model.Focus(parentKeyMap)
	} else {
		return AnnounceKeyMapCmd(parentKeyMap)
	}
}

func (s *Slot) TryBlur() {
	if model, ok := s.Model.(Focusable); ok {
		model.Blur()
	}
}

func BorrowSlotModel[T tea.Model](s *Slot, fn func(m *T)) bool {
	if model, ok := s.Model.(T); ok {
		fn(&model)
		s.Model = model
		return true
	}
	return false
}
