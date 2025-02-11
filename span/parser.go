package span

import (
	"errors"
	"strconv"
	"strings"
)

const (
	TIMESTAMP_SIZE = 5
	RANGE_SIZE     = 13
)

func parseTimestamp(input string) (Timestamp, error) {

	if len(input) != TIMESTAMP_SIZE {
		return Timestamp{}, errors.New("timestamp must be in the format hh:mm")
	}

	parts := strings.Split(input, ":")

	if len(parts) != 2 || len(parts[0]) != 2 || len(parts[1]) != 2 {
		return Timestamp{}, errors.New("timestamp must be in the format hh:mm")
	}

	ts := Timestamp{}
	var err error

	// Parse

	ts.Hour, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return ts, errors.New("unable to parse hours")
	}

	ts.Minute, err = strconv.Atoi(parts[1])
	if err != nil {
		return ts, errors.New("unable to parse minutes")
	}

	// Validate

	if ts.Hour < 0 || ts.Hour > 23 {
		return ts, errors.New("hours must be between 0 and 23")
	}

	if ts.Minute < 0 || ts.Minute > 45 || ts.Minute%15 != 0 {
		return ts, errors.New("minutes must be 0, 15, 30 or 45")
	}

	return ts, nil

}

func parseRange(input string) ([2]Timestamp, []error) {

	rng := [2]Timestamp{}

	if len(input) != RANGE_SIZE {
		return rng, []error{errors.New("range must be in the format hh:mm - hh:mm")}
	}

	errs := make([]error, 0)

	start, err := parseTimestamp(input[0:TIMESTAMP_SIZE])
	if err != nil {
		errs = append(errs, err)
	} else {
		rng[0] = start
	}

	end, err := parseTimestamp(input[RANGE_SIZE-TIMESTAMP_SIZE:])
	if err != nil {
		errs = append(errs, err)
	} else {
		rng[1] = end
	}

	return rng, errs

}

func parseDescriptionAndTag(input string) (string, string) {

	if len(input) == 0 {
		return "", ""
	}

	if !strings.Contains(input, "#") {
		return strings.TrimSpace(input), ""
	}

	h := strings.LastIndex(input, "#")

	desc := strings.TrimSpace(input[0:h])
	tag := strings.TrimSpace(input[h+1:])

	return desc, tag

}

func parseEntry(input string) (Entry, error) {

	if len(input) < RANGE_SIZE {
		return Entry{}, errors.New("unable to parse entry, input too short")
	}

	// @TODO: Change error handling to use logging instead?
	ranges, _ := parseRange(input[0:RANGE_SIZE])
	desc, tag := parseDescriptionAndTag(input[RANGE_SIZE:])

	return newEntry(desc, tag, ranges[0], ranges[1]), nil

}
