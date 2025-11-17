package database

import (
	"worklog/internal/task"
	"worklog/internal/weekRange"
)

func (db *DB) AddTask(t task.Task) (task.Task, error) {
	tasks, err := db.readDBFile()
	if err != nil {
		return task.Task{}, err
	}

	t.ID = len(tasks) + 1

	if err := db.appendDBFile(t); err != nil {
		return task.Task{}, err
	}

	return t, nil
}

func (db *DB) RemoveAllTasks() error {
	var tasks []task.Task
	if err := db.saveDBFile(tasks); err != nil {
		return err
	}

	return nil
}

func (db *DB) RemoveTask(taskId int) error {
	tasks, err := db.readDBFile()

	if err != nil {
		return err
	}

	if len(tasks) < taskId {
		return ErrorNoTask
	}

	tasks = append(tasks[:taskId], tasks[taskId+1:]...)

	if err := db.saveDBFile(tasks); err != nil {
		return err
	}

	return nil
}

func (db *DB) UpdateTask(t task.Task) (task.Task, error) {
	tasks, err := db.readDBFile()
	if err != nil {
		return task.Task{}, err
	}

	tasks[t.ID] = t

	if err := db.saveDBFile(tasks); err != nil {
		return task.Task{}, err
	}

	return t, nil
}

func (db *DB) FindTaskById(taskId int) (task.Task, error) {
	var foundTask task.Task

	tasks, err := db.readDBFile()
	if err != nil {
		return foundTask, err
	}

	if len(tasks) < taskId {
		return foundTask, ErrorNoTask
	}

	foundTask = tasks[taskId]

	return foundTask, nil
}

func (db *DB) FindTasksByName(query string) ([]task.Task, error) {
	var foundTasks []task.Task

	tasks, err := db.readDBFile()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if ok := t.MatchName(query); ok {
			foundTasks = append(foundTasks, t)
		}
	}

	if len(foundTasks) == 0 {
		return nil, ErrorNoTasks
	}

	return foundTasks, nil
}

func (db *DB) FindTasksForWeek(weekRange weekRange.WeekRange) ([]task.Task, error) {
	var foundTasks []task.Task
	tasks, err := db.readDBFile()

	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if ok := weekRange.IsDateInRange(task.Created); ok {
			foundTasks = append(foundTasks, task)
		}
	}

	if len(foundTasks) == 0 {
		return nil, ErrorNoTasks
	}

	return foundTasks, nil
}
