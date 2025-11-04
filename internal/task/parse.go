package task

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StrToTask(text string) (Task, error) {
	var (
		task   Task
		branch string
		repo   string
		tags   []string
	)

	id := text[0:3]
	created := text[6:25]
	text = text[26:]

	// find branch part
	branchPartIndex := strings.Index(text, "@branch:")
	fmt.Printf("branchPartIndex, [%v]\n", branchPartIndex)
	if branchPartIndex != -1 {
		branch = strings.TrimSpace(
			strings.Replace(text[branchPartIndex:], "@branch:", "", 1),
		)
		text = text[:branchPartIndex]
	}

	// find repo part
	repoPartIndex := strings.Index(text, "@repo:")
	if repoPartIndex != -1 {
		repo = strings.TrimSpace(
			strings.Replace(text[repoPartIndex:], "@repo:", "", 1),
		)
		text = text[:repoPartIndex]
	}

	// find tags part
	tagsPartIndex := strings.Index(text, "@tags:")
	if tagsPartIndex != -1 {
		tags = strings.Split(
			strings.TrimSpace(strings.Replace(text[tagsPartIndex:], "@tags:", "", 1)),
			",",
		)

		text = text[:tagsPartIndex]
	}

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
		Name:    text,
		Tags:    tags,
		Branch:  branch,
		Repo:    repo,
	}

	return task, nil
}
