package main

import (
	"github.com/gdamore/tcell"
)

type State struct {
	x, y int
}

func (s *State) MoveLeft() {
	s.x--
}

func (s *State) MoveRight() {
	s.x++
}

func (s *State) MoveUp() {
	s.y--
}

func (s *State) MoveDown() {
	s.y++
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
	state := State{}

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
	case tcell.KeyLeft:
		state.MoveLeft()
	case tcell.KeyRight:
		state.MoveRight()
	case tcell.KeyDown:
		state.MoveDown()
	case tcell.KeyUp:
		state.MoveUp()
	}
	drawAndShow(screen, state)
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
		x := state.x + i
		y := state.y
		screen.SetContent(x, y, cellRunes[0], cellRunes[1:], tcell.StyleDefault)
	}
	screen.Show()
}
