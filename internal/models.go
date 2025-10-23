package internal

import "time"

type Tag string

type Task struct {
	id   string
	date time.Time
	name string
	tags []Tag
}

func NewTask(name string, date time.Time, tags []Tag) Task {
	return Task{
		id:   "1",
		name: name,
		date: date,
		tags: tags,
	}
}
