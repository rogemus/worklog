package internal

import "errors"

var ErrorCannotCreateDB = errors.New("Cannot create DB")
var ErrorCannotReadDB = errors.New("Cannot read data from DB")
var ErrorCannotUpdateDB = errors.New("Cannot update data in the DB")
var ErrorCannotParse = errors.New("Cannot parse DB file")
var ErrorCannotStringify = errors.New("Cannot convert DB struct to string")

var ErrorNoTask = errors.New("No task found")
var ErrorNoTasks = errors.New("No tasks found")
var ErrorNoTaskToUpdate = errors.New("No task to update")
