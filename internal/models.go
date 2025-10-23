package internal

import (
	"strings"
	"time"
)

type Tag string

type Task struct {
	Id   int16     `json:"id"`
	Date time.Time `json:"data"`
	Name string    `json:"name"`
	Tags []Tag     `json:"tags"`
}

func (t Task) MatchId(id int16) bool {
	if t.Id == id {
		return true
	}

	return false
}

func (t Task) InRange(weekRange WeekRange) bool {
	if t.Date.After(weekRange.Start) && t.Date.Before(weekRange.End) {
		return true
	}

	return false
}

func (t Task) MatchName(query string) bool {
	if strings.Contains(t.Name, query) {
		return true
	}

	return false
}

func NewTask(name string, date time.Time, tags []Tag) Task {
	return Task{
		Id:   -1,
		Name: name,
		Date: date,
		Tags: []Tag{},
	}
}

type DBSchema struct {
	LastIndex int16  `json:"lastIndex"`
	Tasks     []Task `json:"tasks"`
}

type WeekRange struct {
	Start time.Time
	End   time.Time
}

func FirstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, 7)
		isoYear, isoWeek = date.ISOWeek()
	}

	return date
}

func (w *WeekRange) CalculateWeekRange(date time.Time) {
	// NOTE: go back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
	}

	w.Start = date
	w.End = date.AddDate(0, 0, 6)
}

func NewWeekRange(date time.Time) WeekRange {
	var week WeekRange
	week.CalculateWeekRange(date)

	return week
}
