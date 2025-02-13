package span

type Timestamp struct {
	Hour, Minute int
}

type Entry struct {
	Warnings         []string
	Description, Tag string
	Start, End       Timestamp
	Duration         int
}

func newEntry(desc, tag string, start, end Timestamp, warnings []string) Entry {

	entry := Entry{
		Warnings:    warnings,
		Description: desc,
		Tag:         tag,
		Start:       start,
		End:         end,
	}

	entry.CalculateDuration()

	return entry

}

func (e *Entry) CalculateDuration() {

	dur := (e.End.Hour - e.Start.Hour) * 60
	if dur < 0 {
		e.Warnings = append(e.Warnings, "negative duration")
		return
	}

	dur += e.End.Minute - e.Start.Minute
	if dur < 0 {
		e.Warnings = append(e.Warnings, "negative duration")
		return
	}

	e.Duration = dur

}
