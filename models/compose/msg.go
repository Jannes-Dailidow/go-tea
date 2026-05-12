package compose

import (
	"github.com/jannes-dailidow/go-tea/util"
)

type ChangeFocusMsg struct {
	id    util.ModelId
	focus Focus
}
