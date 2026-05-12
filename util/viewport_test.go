package util_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/jannes-dailidow/go-slicest"
	"github.com/jannes-dailidow/go-tea/util"
)

func joinLines(lines []string) string   { return strings.Join(lines, "\n") }
func concatenate(lines []string) string { return strings.Join(lines, "") }

var testRows = [][]string{
	{"R1C1", "R1C2", "R1C3", "R1C4", "R1C5"},
	{"R2C1", "R2C2", "R2C3", "R2C4", "R2C5"},
	{"R3C1", "R3C2", "R3C3", "R3C4", "R3C5"},
	{"R4C1", "R4C2", "R4C3", "R4C4", "R4C5"},
	{"R5C1", "R5C2", "R5C3", "R5C4", "R5C5"},
}
var testContent = joinLines(slicest.Map(testRows, concatenate))

func testSubset(from, to util.Vec2[int]) string {
	strs := make([]string, 0, to.Y-from.Y+1)
	for i := from.Y - 1; i <= to.Y-1; i++ {
		strs = append(strs, concatenate(testRows[i][from.X-1:to.X]))
	}
	return joinLines(strs)
}
func newSize(x, y int) util.Size     { return util.Size{x * 4, y} }
func newPos(x, y int) util.Vec2[int] { return util.Vec2[int]{(x - 1) * 4, y - 1} }

func TestRenderContentInViewportAlignXY(t *testing.T) {
	tests := []struct {
		name  string
		vSize util.Size
		tPos  util.Vec2[int]
		tSize util.Size
		align util.Vec2[lipgloss.Position]
		want  string
	}{
		{
			name:  "1x1 cell",
			vSize: newSize(1, 1),
			tPos:  newPos(2, 2),
			tSize: newSize(1, 1),
			align: util.Vec2[lipgloss.Position]{lipgloss.Center, lipgloss.Center},
			want:  testSubset(util.Vec2[int]{2, 2}, util.Vec2[int]{2, 2}),
		},
		{
			name:  "inner 3x3 Top Left",
			vSize: newSize(3, 3),
			tPos:  newPos(3, 3),
			tSize: newSize(1, 1),
			align: util.Vec2[lipgloss.Position]{lipgloss.Top, lipgloss.Left},
			want:  testSubset(util.Vec2[int]{3, 3}, util.Vec2[int]{5, 5}),
		},
		{
			name:  "inner 3x3 Center Bottom",
			vSize: newSize(3, 3),
			tPos:  newPos(3, 3),
			tSize: newSize(1, 1),
			align: util.Vec2[lipgloss.Position]{lipgloss.Center, lipgloss.Bottom},
			want:  testSubset(util.Vec2[int]{2, 1}, util.Vec2[int]{4, 3}),
		},
		{
			name:  "inner 3x3 Right Center",
			vSize: newSize(3, 3),
			tPos:  newPos(3, 3),
			tSize: newSize(1, 1),
			align: util.Vec2[lipgloss.Position]{lipgloss.Right, lipgloss.Center},
			want:  testSubset(util.Vec2[int]{1, 2}, util.Vec2[int]{3, 4}),
		},
		{
			name:  "inner 3x3 Center Center",
			vSize: newSize(3, 3),
			tPos:  newPos(3, 3),
			tSize: newSize(1, 1),
			align: util.Vec2[lipgloss.Position]{lipgloss.Center, lipgloss.Center},
			want:  testSubset(util.Vec2[int]{2, 2}, util.Vec2[int]{4, 4}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.RenderContentInViewportAlign(testContent, tt.vSize, tt.tPos, tt.tSize, tt.align)
			if got != tt.want {
				t.Errorf("got:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}
