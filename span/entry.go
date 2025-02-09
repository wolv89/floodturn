package span

type Timestamp struct {
	Hour, Minute int
}

type Entry struct {
	Description, Tag string
	Start, End       Timestamp
	Duration         int
}
