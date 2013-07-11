package main

import (
	"github.com/nsf/termbox-go"
)

type confirmExitState struct {
	p state
}

func (s *confirmExitState) parent() state {
	return s.p
}

func (s *confirmExitState) handle(event termbox.Event) bool {
	return true
}

var confirmExitStateText = [][]rune{
	[]rune("Press ESC to go back to the main menu."),
	[]rune("Press any other key to exit."),
}

func (s *confirmExitState) paint(w, h int) {
	left, top := lefttop(w, h, confirmExitStateText)

	for i, line := range confirmExitStateText {
		drawText(left, top+i, line, white, black)
	}
}

type mainMenuState struct {
	selected int
}

func (s *mainMenuState) parent() state {
	return &confirmExitState{s}
}

func (s *mainMenuState) handle(event termbox.Event) bool {
	switch event.Key {
	case termbox.KeyArrowDown:
		s.selected = (s.selected + 1) % len(mainMenuStateText)
		repaint()

	case termbox.KeyArrowUp:
		s.selected = (s.selected + len(mainMenuStateText) - 1) % len(mainMenuStateText)
		repaint()
	}
	return false
}

var mainMenuStateText = [][]rune{
	[]rune("Start Fresh"),
	[]rune("Resume Game"),
}

var mainMenuBackHint = []rune("Push ESC to go back to the previous interface.")

func (s *mainMenuState) paint(w, h int) {
	left, top := lefttop(w, h, mainMenuStateText)

	for i, line := range mainMenuStateText {
		if i == s.selected {
			for j := 0; j < w; j++ {
				termbox.SetCell(j, top+i, ' ', black, white)
			}
			drawText(left, top+i, line, black|bold, white)
		} else {
			drawText(left, top+i, line, white, black)
		}
	}

	drawText(0, h-1, mainMenuBackHint, white, black)
}
