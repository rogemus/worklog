package internal

import (
	"encoding/json"
	"errors"
	"os"
)

var DB_FILE_PATH = "./db.json"

func CreateDBFile() error {
	f, err := os.Create(DB_FILE_PATH)

	if err != nil {
		return err
	}

	content, err := json.Marshal(DBSchema{})
	if err != nil {
		panic(err)
	}

	_, err = f.Write(content)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	return nil
}

func ReadFile() []Task {
	// NOTE: check if file exist
	if _, err := os.Stat(DB_FILE_PATH); errors.Is(err, os.ErrNotExist) {
		// NOTE: create new db if exist
		if err = CreateDBFile(); err != nil {
			panic(err)
		}

	}

	file, err := os.ReadFile(DB_FILE_PATH)
	if err != nil {
		panic(err)
	}

	var dbData DBSchema

	if err := json.Unmarshal(file, &dbData); err != nil {
		panic(err)
	}

	return dbData.Tasks
}

func UpdateFile(tasks []Task) {
	updatedDB := DBSchema{Tasks: tasks}

	json, err := json.Marshal(updatedDB)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(DB_FILE_PATH, json, 0644); err != nil {
		panic(err)
	}

}
