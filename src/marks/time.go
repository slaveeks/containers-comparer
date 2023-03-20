package marks

import (
	"fmt"
	"time"
)

type Time struct {
	start time.Time
	end   time.Time
}

func (t *Time) TakeDiff() {
	diff := t.end.Sub(t.start)

	fmt.Println(diff)
}

func CreateTimeMark(start time.Time, end time.Time) *Time {
	return &Time{start, end}
}
