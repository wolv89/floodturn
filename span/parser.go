package span

import (
	"errors"
	"strconv"
	"strings"
)

func parseTimestamp(input string) (Timestamp, error) {

	if len(input) != 5 {
		return Timestamp{}, errors.New("timestamp must be in the format hh:mm")
	}

	parts := strings.Split(input, ":")

	if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 2 {
		return Timestamp{}, errors.New("timestamp must be in the format hh:mm")
	}

	ts := Timestamp{}
	var err error

	ts.Hour, err = strconv.Atoi(parts[0])
	if err != nil {
		return ts, errors.New("unable to parse hours")
	}

	ts.Minute, err = strconv.Atoi(parts[1])
	if err != nil {
		return ts, errors.New("unable to parse minutes")
	}

	// @TODO: check ranges, ie hour 0-23, minute 0-59
	// (Or should minute only allow multiples of 15?)

	return ts, nil

}
