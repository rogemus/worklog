package internal

import (
	"fmt"
	"time"
)

type DB struct{}

func NewDb() *DB {
	return &DB{}
}

func (db *DB) AddTask(task Task) {
	fmt.Printf("New task [%s] | Task Date [%s]", task.name, task.date.Format("2006-01-02 15:04:05"))
}

func (db *DB) UpdateTask(task Task) {}

func (db *DB) RemoveTask(id string) {}

func (db *DB) GetTask(id string) Task {
	return Task{}
}

func (db *DB) FindByName(name string) []Task {
	return nil
}

func (db *DB) GetTasks(startDate time.Time) []Task {
	return nil
}
