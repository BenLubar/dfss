package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func pollEvents(events chan termbox.Event) {
	for {
		events <- termbox.PollEvent()
	}
}

type state interface {
	parent() state
	handle(termbox.Event) (exit bool)
	paint(w, h int)
}

var currentState state

func handleEvent(event termbox.Event) bool {
	switch event.Type {
	case termbox.EventError:
		log.Printf("[ERR] termbox error event: %v", event.Err)

	case termbox.EventResize:
		repaint()

	case termbox.EventKey:
		switch event.Key {
		case termbox.KeyEsc:
			currentState = currentState.parent()
			repaint()
			return false

		case termbox.KeyCtrlBackslash:
			panic("stack trace requested")
		}
		return currentState.handle(event)
	}
	return false
}
