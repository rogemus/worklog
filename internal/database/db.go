package database

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/rogemus/worklog/internal/task"
)

type DB struct {
	path     string
	fileName string

	Tasks task.Task
}

const defaultFileName = "worklog.txt"

func NewDB(path string) *DB {
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
	if strings.HasPrefix(db.path, ".config") {
		dirname, err := os.UserHomeDir()

		if err != nil {
			dirname = "."
		}

		return dirname + "/" + db.path + "/" + db.fileName
	}

	return db.path + "/" + db.fileName
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
	scanner.Split(bufio.ScanLines)
	lineNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		task := task.NewTaskFromStr(line, lineNumber)

		if task.Name != "" {
			tasks = append(tasks, task)
			lineNumber = lineNumber + 1
		}
	}

	if err := scanner.Err(); err != nil {
		return tasks, ErrorCannotParse
	}

	return tasks, nil
}

func (db *DB) appendDBFile(task task.Task) error {
	file, err := os.OpenFile(db.getPath(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return ErrorCannotReadDB
	}

	defer file.Close()

	text := task.ToString() + "\n"
	if _, err = file.WriteString(text); err != nil {
		return ErrorCannotUpdateDB
	}

	return nil
}

func (db *DB) saveDBFile(tasks []task.Task) error {
	var tasksStr []string

	for _, task := range tasks {
		tasksStr = append(tasksStr, task.ToString())
	}

	text := ([]byte)(strings.Join(tasksStr, "\n"))
	if err := os.WriteFile(db.getPath(), text, 0644); err != nil {
		return ErrorCannotUpdateDB
	}

	return nil
}
