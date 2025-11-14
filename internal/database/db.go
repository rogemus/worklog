package database

import (
	"bufio"
	"errors"
	"os"
	"worklog/internal/task"
)

type DB struct {
	path     string
	fileName string

	Tasks task.Task
}

const defaultPath = ".config/worklog"
const defaultFileName = "worklog.txt"

func NewDB(path string) *DB {
	if path == "" {
		path = defaultPath
	}

	db := &DB{
		path:     path,
		fileName: defaultFileName,
	}

	err := db.createDBfile()
	if err != nil {
		panic(err)
	}

	return db
}

func (db DB) getPath() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		dirname = "."
	}

	return dirname + "/" + db.path + "/" + db.fileName
}

func (db *DB) createDBdir() error {
	dirname, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	err = os.MkdirAll(dirname+"/"+db.path, 0755)
	if err != nil {
		return ErrorCannotCreateDBdir
	}

	return nil
}

func (db *DB) createDBfile() error {
	if _, err := os.Stat(db.getPath()); errors.Is(err, os.ErrNotExist) {
		err := db.createDBdir()
		if err != nil {
			return ErrorCannotCreateDBdir
		}

		file, err := os.Create(db.getPath())
		if err != nil {
			return ErrorCannotCreateDB
		}
		defer file.Close()
	}

	return nil
}

func (db *DB) readDBFile() ([]task.Task, error) {
	var tasks []task.Task

	file, err := os.Open(db.getPath())
	if err != nil {
		return tasks, ErrorCannotReadDB
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		task := task.NewTaskFromStr(scanner.Text())

		if task.Name != "" {
			tasks = append(tasks, task)
		}
	}

	if err := scanner.Err(); err != nil {
		return tasks, ErrorCannotParse
	}

	return tasks, nil
}

func (db *DB) appendDBFile(task task.Task) error {
	file, err := os.OpenFile(db.getPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return ErrorCannotUpdateDB
	}

	defer file.Close()

	var text string

	text = task.String() + "\n"

	if _, err = file.WriteString(text); err != nil {
		return ErrorCannotUpdateDB
	}

	return nil
}

func (db *DB) saveDBFile(tasks []task.Task) error {
	file, err := os.OpenFile(db.getPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return ErrorCannotUpdateDB
	}

	defer file.Close()

	var text string

	for _, task := range tasks {
		text = text + task.String() + "\n"
	}

	if _, err = file.WriteString(text); err != nil {
		return ErrorCannotUpdateDB
	}

	return nil
}
