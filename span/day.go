package span

import "time"

type Day struct {
	Entries []Entry
	Weekday time.Weekday
}

func newDay(w int) Day {
	return Day{
		Entries: make([]Entry, 0),
		Weekday: time.Weekday(w),
	}
}
