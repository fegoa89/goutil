// Package timespan is a collection of functions to manipulate time intervals
package timespan

import (
	"time"
)

// TimeSpan is the struct that holds the start and end time objects
type TimeSpan struct {
	start, end time.Time
}

// Creates a TimeSpan object that holds the start and end time object
func TimeSpanStruct(s time.Time, e time.Time) TimeSpan {
	return TimeSpan{
		start: s,
		end:   e,
	}
}

// Returns the start time
func (t TimeSpan) Start() time.Time {
	return t.start
}

// Returns the end time
func (t TimeSpan) End() time.Time {
	return t.end
}

// Checks if the start date is older than end date
func (t TimeSpan) ValidDateRange() bool {
	return t.end.After(t.start)
}

// Returns the quantity of days between two time objects
func (t TimeSpan) DaysBetween() int64 {
	return int64(t.end.Sub(t.start).Hours() / 24)
}

// Returns the quantity of hours between two time objects
func (t TimeSpan) HoursBetween() int64 {
	return int64(t.end.Sub(t.start).Hours())
}

// Returns the quantity of minutes between two time objects
func (t TimeSpan) MinutesBetween() int64 {
	return int64(t.end.Sub(t.start).Minutes())
}

// Returns the quantity of seconds between two time objects
func (t TimeSpan) SecondsBetween() int64 {
	return int64(t.end.Sub(t.start).Seconds())
}

// Returns the quantity of months between two time objects
func (t TimeSpan) MonthsBetween() int64 {
	return int64(t.DaysBetween() / 30)
}
