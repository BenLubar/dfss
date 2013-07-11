package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	logFile = flag.String("log", "/dev/null", "file path to use for logging")
)

func main() {
	flag.Parse()

	f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("[SETUP] opening log file: %v", err)
	}
	defer f.Close()

	err = termbox.Init()
	if err != nil {
		log.Fatalf("[SETUP] termbox init failed: %v", err)
	}
	defer termbox.Close()

	log.SetOutput(f)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

	repaint()

	events := make(chan termbox.Event)
	go pollEvents(events)

	currentState = &mainMenuState{}

	ticker := time.Tick(time.Minute / 100)

	for {
		select {
		case <-paintScheduler:
			paint()

		case event := <-events:
			if handleEvent(event) {
				return
			}

		case <-ticker:
			tick()
		}
	}
}
