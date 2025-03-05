package api

import (
	"net/http"
	"time"
)

type Datestamp struct {
	Formal   string `json:"formal"`
	Friendly string `json:"friendly"`
}

const (
	SELECTION_WEEKS = 3
	SELECTION_DAYS  = SELECTION_WEEKS * 7
)

func handleGetDateRange(w http.ResponseWriter, req *http.Request) {

	today := time.Now()
	daterange := getSelectionRange(today)

	jsonResponse(w, http.StatusOK, daterange)

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
