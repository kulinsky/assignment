package assignment

import "time"

type (
	DateStart time.Time
	DateEnd   time.Time
)

func (d DateStart) Time() time.Time {
	return time.Time(d)
}

func (d DateEnd) Time() time.Time {
	return time.Time(d)
}
