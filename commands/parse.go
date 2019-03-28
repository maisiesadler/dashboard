package commands

import (
  "sync"
  "github.com/gdamore/tcell"
)

func Parse(s tcell.Screen) {
  var wg sync.WaitGroup
  wg.Add(1)
  go parse(s, &wg)
  wg.Wait()
}

func parse(s tcell.Screen, wg *sync.WaitGroup) {
  for {
    ev := s.PollEvent()
    switch ev := ev.(type) {
    case *tcell.EventKey:
      switch ev.Key() {
      case tcell.KeyEscape, tcell.KeyCtrlC:
        wg.Done()
        return
      }
    case *tcell.EventResize:
      s.Sync()
    }
  }
}

