package internal

import (
	"fmt"
	"time"
)

type WeekRange struct {
	Start time.Time
	End   time.Time
}

func (w *WeekRange) CalculateWeekRange(date time.Time) {
	// NOTE: go back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 1, time.UTC)

	w.Start = date
	w.End = date.AddDate(0, 0, 6)
}

func (w WeekRange) GetFormatedRange() string {
	startFormatted := w.Start.Format("02-01-2006")
	endFormatted := w.End.Format("02-01-2006")
	return fmt.Sprintf("[%s - %s]", startFormatted, endFormatted)
}

func NewWeekRange(date time.Time) WeekRange {
	var week WeekRange
	week.CalculateWeekRange(date)

	return week
}
