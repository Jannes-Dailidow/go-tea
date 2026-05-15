package main

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/jannes-dailidow/go-tea/models/compose"
	"github.com/jannes-dailidow/go-tea/models/keyhelp"
	"github.com/jannes-dailidow/go-tea/util"
)

func main() {
	tea.NewProgram(
		compose.New(
			compose.WithPrimary(&DummyModel{
				KeyMap: util.KeyBindingList{key.NewBinding(
					key.WithKeys("esc"),
					key.WithHelp("esc", "quit"),
				)},
			}),
			compose.WithSecondary(keyhelp.New()),
			compose.WithFocussed(true),
			compose.WithMsgTypeRoute[util.AnnounceKeyMapMsg](compose.Secondary),
			compose.WithSizeResolver(func(
				size util.Size,
				primary, secondary tea.Model,
				focus compose.Focus,
				orientation compose.Orientation,
			) (util.Size, util.Size) {
				keyhelpHeight := min(size.Height, lipgloss.Height(secondary.View().Content))
				return util.Size{size.Width, size.Height - keyhelpHeight}, util.Size{size.Width, keyhelpHeight}
			}),
		),
	).Run()
}
