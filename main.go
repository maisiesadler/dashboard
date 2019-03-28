package main

import (
  "log"
  "os"
  "regexp"
  "sort"
  "strconv"

  "./commands"
  "./viewmodels"

  "github.com/gdamore/tcell"
  "github.com/gdamore/tcell/views"
  "github.com/maisiesadler/logtail"
  "github.com/maisiesadler/logtail/writers"
)

func main() {
  args := os.Args[1:]
  if len(args) != 2 {
    os.Exit(1)
  }
  file := args[0]
  s := initScreen()
  fw := writers.NewSummaryWriter(regexp.MustCompile(args[1]))
  go func() {
    for range fw.Updates() {
      v := View(parse(fw.Summary.Read()))
      DrawModules(s, []*views.CellView { v })
    }
  }()
  done := logtail.Run(file, fw)
  commands.Parse(s)
  done<-true
}

type byCount []writers.LineAndCount

func (s byCount) Len() int {
    return len(s)
}
func (s byCount) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byCount) Less(i, j int) bool {
    ii := s[i]
    jj := s[j]
    if ii.Count == jj.Count {
      return ii.Line < jj.Line
    }
    return ii.Count > jj.Count
}

func parse(lac []writers.LineAndCount) []string {
  sort.Sort(byCount(lac))
  r := make([]string, len(lac))
  for i, v := range lac {
    r[i] = v.Line + " (" + strconv.Itoa(v.Count) + ")"
  }
  return r
}

func View(d []string) *views.CellView {
  cell := views.NewCellView()
  cell.SetModel(viewmodels.FromList(d))
  cell.Draw()
  return cell
}

func DrawModules(s tcell.Screen, vs []*views.CellView) {
  s.Clear()

  inner := views.NewBoxLayout(views.Horizontal)
  inner.SetView(s)
  for _, v := range vs {
    inner.AddWidget(v, 0.5)
  }
  inner.Draw()
  s.Show()
}

func initScreen() tcell.Screen {
  tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
  s, e := tcell.NewScreen()
  if e != nil {
    log.Printf("%v\n", e)
    os.Exit(1)
  }
  if e = s.Init(); e != nil {
    log.Printf("%v\n", e)
    os.Exit(1)
  }
  s.Clear()
  return s
}

func logger(l string) {
  f, err := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
	log.Println(err)
  }
  defer f.Close()

  logger := log.New(f, "prefix", log.LstdFlags)
  logger.Println(l)
}
