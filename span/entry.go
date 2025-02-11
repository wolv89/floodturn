package span

type Timestamp struct {
	Hour, Minute int
}

type Entry struct {
	Description, Tag string
	Start, End       Timestamp
	Duration         int
}

func newEntry(desc, tag string, start, end Timestamp) Entry {

	entry := Entry{
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
		return
	}

	dur += e.End.Minute - e.Start.Minute
	if dur < 0 {
		return
	}

	e.Duration = dur

}
