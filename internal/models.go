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

func (t Task) MatchId(id int16) (Task, bool) {
	if t.Id == id {
		return t, true
	}

	return Task{}, false
}

func (t Task) MatchName(query string) (Task, bool) {
	if strings.Contains(t.Name, query) {
		return t, true
	}

	return Task{}, false
}

type DBSchema struct {
	LastIndex int16  `json:"lastIndex"`
	Tasks     []Task `json:"tasks"`
}

func NewTask(name string, date time.Time, tags []Tag) Task {
	return Task{
		Id:   -1,
		Name: name,
		Date: date,
		Tags: []Tag{},
	}
}
