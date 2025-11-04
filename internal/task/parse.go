package task

import (
	// "fmt"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

// 001: [21-10-1994 21:37] "XXX-123 this is my task" @tags:feat @issue_id:XXX-123 @repo:<name> @branch:feat/XXX-123-this-is-my-task
func StrToTask(text string) (Task, error) {
	var task Task

	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	s.Filename = "task"

	id := text[0:3]
	created := text[6:25]
	text = text[26:]

	var (
		col int
		// partStart int
		// partEnd   int
	)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// txt := s.TokenText()

		// switch txt {
		// case "[":
		// 	partStart = col
		// case "]":
		// 	created = text[partStart : col-1]
		// 	partStart = 0
		// }
		col++
	}

	// var parts []string
	// // id
	// // created
	// // name
	// // tags
	// // issue_id
	// // repo
	// // branch

	parsedId, err := strconv.ParseInt(id, 10, 10)
	if err != nil {
		return Task{}, ErrorParseInvalidId
	}

	parsedCreated, err := time.Parse(time.DateTime, created)
	if err != nil {
		return Task{}, ErrorParseInvalidCreatedDate
	}

	task = Task{
		ID:      int(parsedId),
		Created: parsedCreated,
	}

	return task, nil
}
