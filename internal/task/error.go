package task

import "errors"

var ErrorParseInvalidCreatedDate = errors.New("Parsing Error: invalid task created date. Need to be in `YYYY-MM-DD HH:mm:ss` format")

var ErrorInvalidDate = errors.New("Invalid task date format, use YYYY-MM-DD format")
var ErrorEmptyName = errors.New("No name provided")
var ErrorNotInGitRepo = errors.New("Command was not run in git repo")
var ErrorMainBranch = errors.New("Cannot use `main`/`master` as task name")

var ErrorFindMultipleFlags = errors.New("Cannot search using multiple flags")
var ErrorFindWithoutTag = errors.New("Tags not provided")
