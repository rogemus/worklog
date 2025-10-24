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

	w.Start = date
	w.End = date.AddDate(0, 0, 6)
}

func (w WeekRange) GetFormatedRange() string {
	startFormatted := w.Start.Format("02-01-2006")
	endFormatted := w.End.Format("02-01-2005")
	return fmt.Sprintf("[%s - %s]", startFormatted, endFormatted)
}

func (w WeekRange) RenderTitle() string {
	title := fmt.Sprintf("Tasks for week: %s", w.GetFormatedRange())

	return listTitleStyles.Render(title)
}

func NewWeekRange(date time.Time) WeekRange {
	var week WeekRange
	week.CalculateWeekRange(date)

	return week
}
