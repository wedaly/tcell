package main

import (
	"github.com/gdamore/tcell"
)

type State struct {
	offset int
}

func (s *State) MoveUp() {
	s.offset--
	if s.offset < 0 {
		s.offset = 0
	}
}

func (s *State) MoveDown() {
	s.offset++
	if s.offset > 10 {
		s.offset = 10
	}
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	} else {
		defer screen.Fini()
	}

	quit := make(chan struct{})
	state := State{offset: 0}

	go func() {
		drawAndShow(screen, &state)
		for {
			event := screen.PollEvent()
			switch event := event.(type) {
			case *tcell.EventKey:
				handleKey(event, screen, &state, quit)

			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	<-quit
}

func handleKey(event *tcell.EventKey, screen tcell.Screen, state *State, quit chan struct{}) {
	switch event.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		close(quit)
		return
	case tcell.KeyCtrlL:
		screen.Sync()
	case tcell.KeyEnter:
		drawAndShow(screen, state)
	case tcell.KeyDown:
		state.MoveDown()
		drawAndShow(screen, state)
	case tcell.KeyUp:
		state.MoveUp()
		drawAndShow(screen, state)
	}
}

func drawAndShow(screen tcell.Screen, state *State) {
	content := [][]rune{
		{3588, 3657, 3635},
		{3594, 3641},
		{3585, 3641, 3657},
		{3610},
		{3619},
		{3619},
		{3621, 3633},
		{3591},
		{3585, 3660},
		{32},
		{3631},
		{10},
	}

	screen.Clear()
	for i, cellRunes := range content {
		screen.SetContent(i, state.offset, cellRunes[0], cellRunes[1:], tcell.StyleDefault)
	}
	screen.Show()
}
