package timespan

import (
	"testing"
	"time"
)

var (
	layout          string = "2006 01 02 15 04"
	startDateString string = "2015 11 11 00 00"
	endDateString   string = "2015 11 11 16 00"
)

func TestStart(t *testing.T) {
	timeSpanStart := timeSpanStruct().Start()
	startTime := startTimeObject(startDateString)
	if timeSpanStart != startTime {
		t.Errorf("Start() returns %q, want %q", timeSpanStart, startTime)
	}
}

func TestEnd(t *testing.T) {
	timeSpanEnd := timeSpanStruct().End()
	endTime := endTimeObject(endDateString)
	if timeSpanEnd != endTime {
		t.Errorf("End() returns %q, want %q", timeSpanEnd, endTime)
	}
}

func TestValidDateRange(t *testing.T) {
	cases := []struct {
		start, end       string
		expectedResponse bool
	}{
		{"2018 01 11 00 00", "2018 10 11 16 00", true},
		{"2018 05 11 00 00", "2018 04 11 16 00", false},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).ValidDateRange()

		if got != c.expectedResponse {
			t.Errorf("ValidDateRange() returns %q, want %q", got, c.expectedResponse)
		}
	}
}

func TestDaysBetween(t *testing.T) {
	cases := []struct {
		start, end string
		days       int64
	}{
		{"2018 01 11 00 00", "2018 10 11 00 00", 273},
		{"2018 01 11 00 00", "2018 01 12 00 00", 1},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).DaysBetween()

		if got != c.days {
			t.Errorf("DaysBetween() returns %q, want %q", got, c.days)
		}
	}
}

func TestHoursBetween(t *testing.T) {
	cases := []struct {
		start, end string
		hours      int64
	}{
		{"2018 10 11 00 00", "2018 10 11 03 00", 3},
		{"2018 10 11 00 00", "2018 10 12 00 00", 24},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).HoursBetween()

		if got != c.hours {
			t.Errorf("HoursBetween() returns %q, want %q", got, c.hours)
		}
	}
}

func TestSecondsBetween(t *testing.T) {
	cases := []struct {
		start, end string
		seconds    int64
	}{
		{"2018 10 11 00 00", "2018 10 11 00 03", 180},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).SecondsBetween()

		if got != c.seconds {
			t.Errorf("SecondsBetween() returns %q, want %q", got, c.seconds)
		}
	}
}

func TestMonthsBetween(t *testing.T) {
	cases := []struct {
		start, end string
		seconds    int64
	}{
		{"2018 10 01 00 00", "2018 11 01 00 00", 1},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).MonthsBetween()

		if got != c.seconds {
			t.Errorf("MonthsBetween() returns %q, want %q", got, c.seconds)
		}
	}
}

func TestMinutesBetween(t *testing.T) {
	cases := []struct {
		start, end string
		minutes    int64
	}{
		{"2018 10 01 00 00", "2018 10 01 00 01", 1},
	}

	for _, c := range cases {
		got := TimeSpanStruct(startTimeObject(c.start), endTimeObject(c.end)).MinutesBetween()

		if got != c.minutes {
			t.Errorf("MinutesBetween() returns %q, want %q", got, c.minutes)
		}
	}
}

func timeSpanStruct() TimeSpan {
	return TimeSpanStruct(startTimeObject(startDateString), endTimeObject(endDateString))
}

func startTimeObject(dateString string) time.Time {
	t, _ := time.Parse(layout, dateString)
	return t
}

func endTimeObject(dateString string) time.Time {
	t, _ := time.Parse(layout, dateString)
	return t
}
