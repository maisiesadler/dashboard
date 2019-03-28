package viewmodels

import (
  "github.com/gdamore/tcell"
)

type SimpleListState struct {
  folders        []string
  xbound, ybound int
}

func EmptyState() *SimpleListState {
  return &SimpleListState{[]string{}, 0, 0}
}

func FromList(list []string) *SimpleListState {
  lines := list
  xbound := 0
  ybound := len(lines)
  for index, line := range lines {
    if len(line)-1 > xbound {
      xbound = len(line) - 1
    }
    lines[index] = line
  }
  return &SimpleListState{lines, xbound, ybound}
}

func (vs SimpleListState) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
  style := tcell.StyleDefault
  if y < len(vs.folders) {
    line := vs.folders[y]
    if x < len(vs.folders[y]) {
      return rune(line[x]), style, nil, 1
    }
  }
  return ' ', style, nil, 1
}
func (vs SimpleListState) GetBounds() (int, int) {
  return vs.xbound, vs.ybound
}
func (SimpleListState) SetCursor(int, int) {
}

func (SimpleListState) GetCursor() (int, int, bool, bool) {
  return 0, 0, false, false
}
func (SimpleListState) MoveCursor(offx, offy int) {

}
