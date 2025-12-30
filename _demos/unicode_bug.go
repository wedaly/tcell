//go:build ignore
// +build ignore

// Copyright 2025 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

const unicodeStr = "😶‍🌫️ E13.1 face in clouds"

func draw(s tcell.Screen, col int, row int) {
	s.Clear()
	s.PutStrStyled(col, row, unicodeStr, tcell.StyleDefault)
	s.Show()
}

// This program demonstrates a bug in unicode rendering.
// Use the arrow keys to move the text around. Each redraw should clear
// the screen and draw the text at the new position, but instead there's a
// "trail" of dirty cells after the cloud emoji.
func main() {

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	col, row := 0, 0

	draw(s, col, row)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			draw(s, col, row)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				s.Fini()
				os.Exit(0)
			case tcell.KeyUp:
				row--
				draw(s, col, row)
			case tcell.KeyDown:
				row++
				draw(s, col, row)
			case tcell.KeyLeft:
				col--
				draw(s, col, row)
			case tcell.KeyRight:
				col++
				draw(s, col, row)
			}
		}
	}
}
