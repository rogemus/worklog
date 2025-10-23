package internal

import "time"

type Tag string

type Task struct {
	Id   string    `json:"id"`
	Date time.Time `json:"data"`
	Name string    `json:"name"`
	Tags []Tag     `json:"tags"`
}

type DBSchema struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(name string, date time.Time, tags []Tag) Task {
	return Task{
		Id:   "1",
		Name: name,
		Date: date,
		Tags: []Tag{},
	}
}
