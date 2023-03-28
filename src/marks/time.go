package marks

import (
	"fmt"
	"time"
)

type Time struct {
	start time.Time
	end   time.Time
}

func (t *Time) TakeDiff() int {
	diff := t.end.Sub(t.start)

	fmt.Println(diff)
	
	return int(diff.Milliseconds())
}

func CreateTimeMark(start time.Time, end time.Time) *Time {
	return &Time{start, end}
}
