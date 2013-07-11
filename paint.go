package main

import (
	"github.com/nsf/termbox-go"
)

var paintScheduler = make(chan struct{}, 1)

func repaint() {
	select {
	case paintScheduler <- struct{}{}:
	default:
	}
}

func lefttop(w, h int, text [][]rune) (left, top int) {
	left = w / 2
	for _, l := range text {
		if left2 := (w-len(l))/2 - 1; left > left2 {
			left = left2
		}
	}

	top = (h-len(text))/2 - 1

	return
}

func drawText(x, y int, text []rune, fg, bg termbox.Attribute) {
	for i, r := range text {
		termbox.SetCell(x+i, y, r, fg, bg)
	}
}

func paint() {
	termbox.Clear(white, black)
	w, h := termbox.Size()
	if w < 80 {
		w = 80
	}
	if h < 22 {
		h = 22
	}
	currentState.paint(w, h)
	termbox.Flush()
}

const (
	bold  = termbox.AttrBold
	white = termbox.ColorWhite
	black = termbox.ColorBlack
)
