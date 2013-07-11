package main

var tickSchedule [][]func()

func tick() {
	if len(tickSchedule) == 0 {
		return
	}
	scheduled := tickSchedule[0]
	tickSchedule = tickSchedule[1:]
	for _, f := range scheduled {
		f()
	}
}

func schedule(f func()) {
	scheduleAhead(f, 0)
}

func scheduleAhead(f func(), skip int) {
	if len(tickSchedule) <= skip {
		tickSchedule = append(tickSchedule, make([][]func(), skip-len(tickSchedule)+1)...)
	}
	tickSchedule[skip] = append(tickSchedule[skip], f)
}
