package internal

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

var DB_FILE_PATH = "./db.json"

type DB struct {
	path string
}

func NewDb() *DB {
	db := &DB{
		path: DB_FILE_PATH,
	}

	err := db.createDBfile()
	if err != nil {
		panic(err)
	}

	return db
}

func (db *DB) AddTask(task Task) (Task, error) {
	schema, err := db.readDBFile()
	if err != nil {
		return Task{}, err
	}

	newIndex := schema.LastIndex + 1
	task.Id = newIndex
	schema.LastIndex = newIndex
	schema.Tasks = append([]Task{task}, schema.Tasks...)

	if err := db.saveDBFile(schema); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (db *DB) RemoveAllTasks() error {
	schema := DBSchema{LastIndex: 0, Tasks: []Task{}}
	if err := db.saveDBFile(schema); err != nil {
		return err
	}

	return nil
}

func (db *DB) RemoveTask(taskId int16) error {
	schema, err := db.readDBFile()
	if err != nil {
		return err
	}

	idx := slices.IndexFunc(schema.Tasks, func(t Task) bool {
		return t.Id == taskId
	})

	if idx < 0 {
		return ErrorNoTask
	}

	tasks := append(schema.Tasks[:idx], schema.Tasks[idx+1:]...)
	schema.Tasks = tasks

	if err := db.saveDBFile(schema); err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTask(task Task) (Task, error) {
	schema, err := db.readDBFile()
	if err != nil {
		return Task{}, err
	}

	idx := slices.IndexFunc(schema.Tasks, func(t Task) bool {
		return t.Id == task.Id
	})

	if idx < 0 {
		return Task{}, ErrorNoTask
	}

	schema.Tasks[idx] = task

	if err := db.saveDBFile(schema); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (db *DB) FindTaskById(taskId int16) (Task, error) {
	var task Task

	schema, err := db.readDBFile()
	if err != nil {
		return task, err
	}

	idx := slices.IndexFunc(schema.Tasks, func(t Task) bool {
		return t.Id == taskId
	})

	if idx == -1 {
		return task, ErrorNoTask
	}

	task = schema.Tasks[idx]

	return task, nil
}

func (db *DB) FindTasksByName(query string) ([]Task, error) {
	var tasks []Task

	schema, err := db.readDBFile()
	if err != nil {
		return nil, err
	}

	for _, task := range schema.Tasks {
		if ok := task.MatchName(query); ok {
			tasks = append(tasks, task)
		}
	}

	if len(tasks) == 0 {
		return nil, ErrorNoTasks
	}

	return tasks, nil
}

func (db *DB) FindTasksForWeek(weekRange WeekRange) ([]Task, error) {
	var tasks []Task
	schema, err := db.readDBFile()
	if err != nil {
		return nil, err
	}

	for _, task := range schema.Tasks {
		if ok := task.InRange(weekRange); ok {
			tasks = append(tasks, task)
		}
	}

	if len(tasks) == 0 {
		return nil, ErrorNoTasks
	}

	return tasks, nil
}

func (db *DB) createDBfile() error {
	schema := DBSchema{LastIndex: 0, Tasks: []Task{}}

	if _, err := os.Stat(db.path); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(db.path)
		if err != nil {
			return ErrorCannotCreateDB
		}
		defer file.Close()

		json, _ := json.Marshal(schema)

		_, err = file.Write(json)
		if err != nil {
			return ErrorCannotCreateDB
		}
	}

	return nil
}

func (db *DB) readDBFile() (DBSchema, error) {
	var schema DBSchema

	file, err := os.ReadFile(db.path)
	if err != nil {
		return schema, ErrorCannotReadDB
	}

	err = json.Unmarshal(file, &schema)
	if err != nil {
		return schema, ErrorCannotParse
	}

	return schema, nil
}

func (db *DB) saveDBFile(schema DBSchema) error {
	json, err := json.Marshal(schema)
	if err != nil {
		return ErrorCannotStringify
	}

	err = os.WriteFile(db.path, json, 0644)
	if err != nil {
		return ErrorCannotUpdateDB
	}

	return nil
}
