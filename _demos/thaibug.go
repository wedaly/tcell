package main

import (
	"github.com/gdamore/tcell"
)

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

	go func() {
		drawAndShow(screen)
		for {
			event := screen.PollEvent()
			switch event := event.(type) {
			case *tcell.EventKey:
				handleKey(event, screen, quit)

			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	<-quit
}

func handleKey(event *tcell.EventKey, screen tcell.Screen, quit chan struct{}) {
	switch event.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		close(quit)
		return
	case tcell.KeyCtrlL:
		screen.Sync()
	case tcell.KeyEnter:
		drawAndShow(screen)
	}
}

func drawAndShow(screen tcell.Screen) {
}
