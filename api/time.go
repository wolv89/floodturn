package api

import (
	"fmt"
	"time"
)

type Datestamp struct {
	Formal, Friendly string
}

const (
	SELECTION_WEEKS = 3
	SELECTION_DAYS  = SELECTION_WEEKS * 7
)

func Run() {

	today := time.Now()
	quicktest := getSelectionRange(today)

	for _, qt := range quicktest {
		fmt.Println(qt)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

}

func getSelectionRange(day time.Time) []Datestamp {

	sr := make([]Datestamp, SELECTION_DAYS)

	wday := day.Weekday()
	if wday != time.Monday {
		offset := time.Duration(wday) - 1
		if offset == 0 {
			offset = 6
		}
		day = day.Add(time.Hour * -24 * offset)
	}
	day = day.Add(time.Hour * 24 * -7)

	for i := range sr {
		sr[i] = Datestamp{
			Formal:   day.Format(time.DateOnly),
			Friendly: day.Format("Jan 2"),
		}
		day = day.Add(time.Hour * 24)
	}

	return sr

}
