package util

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/jannes-dailidow/go-slicest"
)

type AnnounceKeyMapMsg struct {
	KeyMap help.KeyMap
}

func AnnounceKeyMapCmd(keyMaps ...help.KeyMap) tea.Cmd {
	return TeaMsgToCmd(AnnounceKeyMapMsg{KeyMap: MergeKeyMaps(keyMaps...)})
}

func MergeKeyMaps(keyMaps ...help.KeyMap) help.KeyMap {
	if len(keyMaps) == 1 {
		return keyMaps[0]
	}
	return MergedKeyMaps{KeyMaps: keyMaps}
}

// --- MergedKeyMaps ---

type MergedKeyMaps struct {
	KeyMaps []help.KeyMap
}

// *[MergedKeyMaps] implements [help.KeyMap]
var _ help.KeyMap = (*MergedKeyMaps)(nil)

func (m MergedKeyMaps) ShortHelp() []key.Binding {
	bindings := slicest.Map(m.KeyMaps, func(k help.KeyMap) []key.Binding {
		if k != nil {
			return k.ShortHelp()
		}
		return nil
	})
	return slicest.Flatten(bindings)
}

func (m MergedKeyMaps) FullHelp() [][]key.Binding {
	groups := slicest.Map(m.KeyMaps, func(k help.KeyMap) [][]key.Binding {
		if k != nil {
			return k.FullHelp()
		}
		return nil
	})
	return slicest.Flatten(groups)
}

// --- KeyBindingList ---

type KeyBindingList []key.Binding

// *[KeyBindingList] implements [help.KeyMap]
var _ help.KeyMap = (*KeyBindingList)(nil)

func (km KeyBindingList) ShortHelp() []key.Binding {
	return km
}

func (km KeyBindingList) FullHelp() [][]key.Binding {
	return [][]key.Binding{km}
}
