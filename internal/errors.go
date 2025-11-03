package internal

import "errors"

var ErrorCannotCreateDBdir = errors.New("Cannot create DB dir in user home dir")
var ErrorCannotCreateDB = errors.New("Cannot create DB")
var ErrorCannotReadDB = errors.New("Cannot read data from DB")
var ErrorCannotUpdateDB = errors.New("Cannot update data in the DB")
var ErrorCannotParse = errors.New("Cannot parse DB file")
var ErrorCannotStringify = errors.New("Cannot convert DB struct to string")

var ErrorNoTask = errors.New("No task found")
var ErrorNoTasks = errors.New("No tasks found")
var ErrorNoTaskToUpdate = errors.New("No task to update")
var ErrorInvalidDate = errors.New("Invalid task date format, use YYYY-MM-DD format")
var ErrorEmptyName = errors.New("No name provided")

var ErrorNotInGitRepo = errors.New("Command was not run in git repo")
var ErrorMainBranch = errors.New("Cannot use `main`/`master` as task name")
