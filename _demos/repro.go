package main

import (
	"time"
	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	s.Clear()
	s.Show()

	style := tcell.StyleDefault.
		Background(tcell.ColorMaroon).
		Foreground(tcell.ColorWhite).Bold(true)

	s.SetContent(0, 0, 'a', nil, style)
	s.SetContent(1, 0, 'b', nil, style)
	s.SetContent(2, 0, 'c', nil, style)
	s.Show()

	time.Sleep(time.Second)

	// In tcell v2.4.0, this clears the screen back to its default color.
	// In tcell v2.5.0, this clears the screen to maroon.
	s.Clear()
	s.Show()

	time.Sleep(time.Second * 5)

	s.Fini()
}
