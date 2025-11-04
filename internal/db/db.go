package db

type DB struct {
	path     string
	fileName string
}

const defaultPath = ".config/worklog"
const defaultFileName = "worklog.txt"

func NewDB(path string) *DB {
	if path == "" {
		path = defaultPath
	}

	return &DB{
		path:     path,
		fileName: defaultFileName,
	}
}

func (db *DB) createDBdir() error {
	return nil
}

func (db *DB) createDBfile() error {
	return nil
}

func (db *DB) readDBFile() error {
	return nil
}

func (db *DB) saveDBFile() error {
	return nil
}
