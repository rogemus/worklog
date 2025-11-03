package internal

import (
	"fmt"
	"strconv"
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

func (t Task) RenderWithDate() string {
	idWithStyles := dimStyles.Render(fmt.Sprintf("[%s]", strconv.Itoa(int(t.Id))))
	titleWithStyles := titleStyles.Render(t.Name)
	itemWithStyles := fmt.Sprintf("%s | %s | %s", idWithStyles, t.GetDate(), titleWithStyles)

	return listItemStyles.Render(itemWithStyles)
}

func (t Task) RenderWithHours() string {
	idWithStyles := dimStyles.Render(fmt.Sprintf("[%s]", strconv.Itoa(int(t.Id))))
	titleWithStyles := titleStyles.Render(t.Name)
	itemWithStyles := fmt.Sprintf("%s | %s |%s", idWithStyles, dimStyles.Render(t.GetHour()), tit)

	return listItemStyles.Render(itemWithStyles)
}

func (t Task) Render() string {
	idWithStyles := dimStyles.Render(fmt.Sprintf("[%s]", strconv.Itoa(int(t.Id))))
	titleWithStyles := titleStyles.Render(t.Name)
	itemWithStyles := fmt.Sprintf("%s%s", idWithStyles, titleWithStyles)

	return listItemStyles.Render(itemWithStyles)
}

func (t Task) GetDate() string {
	return t.Date.Format("02-01-2006")
}

func (t Task) GetHour() string {
	return fmt.Sprintf("%d:%d", t.Date.Hour(), t.Date.Minute())
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
