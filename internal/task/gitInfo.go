package task

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GitToTask(branch, repoName string) Task {
	var (
		task    Task
		name    string
		issueId string
	)

	task.Branch = branch
	task.Repo = repoName

	parts := strings.Split(branch, "/")

	if len(parts) > 1 {
		task.Tags = append(task.Tags, strings.ToLower(parts[0]))
		parts := parts[1:]
		branch = strings.Join(parts, "/")
	}

	issuePattern, _ := regexp.Compile("(^[a-zA-Z]+-[0-9]+)|(^[0-9]+)")
	issueId = issuePattern.FindString(branch)

	if issueId != "" {
		task.IssueId = issueId
		branch = strings.Replace(branch, fmt.Sprintf("%s-", issueId), "", 1)
	}

	c := cases.Title(language.English)
	name = strings.ReplaceAll(branch, "-", " ")
	name = c.String(name)

	return task
}
