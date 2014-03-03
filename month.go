package recurrence

import "time"

// A Month represents a month of the year. Just like time.Month.
type Month time.Month

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) IsOccurring(t time.Time) bool {
	return t.Month() == time.Month(m)
}

func (m Month) Occurrences(t TimeRange) chan time.Time {
	return t.occurrencesOfSchedule(m)
}
