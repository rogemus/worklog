package task

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID      int
	Created time.Time
	Name    string
	IssueId string   // optional
	Repo    string   // optional
	Branch  string   // optional
	Tags    []string // optional
}

func (t Task) GetFormatted() string {

	formatted := fmt.Sprintf(
		"%s: [%s] %s",
		t.GetFormattedId(),
		t.GetFormattedCreated(),
		t.Name,
	)

	tags := t.GetFormattedTags()
	issueId := t.GetFormattedIssueId()
	repo := t.GetFormattedRepo()
	branch := t.GetFormattedBranch()

	if tags != "" {
		formatted = formatted + " " + tags
	}

	if issueId != "" {
		formatted = formatted + " " + issueId
	}

	if repo != "" {
		formatted = formatted + " " + repo
	}

	if branch != "" {
		formatted = formatted + " " + branch
	}

	return formatted
}

func (t Task) GetFormattedTags() string {
	if len(t.Tags) == 0 {
		return ""
	}

	return fmt.Sprintf("@tags:%s", strings.Join(t.Tags, ","))
}

func (t Task) GetFormattedId() string {
	return fmt.Sprintf("%03d", t.ID)
}

func (t Task) GetFormattedCreated() string {
	return t.Created.Format(time.DateTime)
}

func (t Task) GetFormattedIssueId() string {
	if t.IssueId == "" {
		return ""
	}

	return fmt.Sprintf("@issue_id:%s", t.IssueId)
}

func (t Task) GetFormattedBranch() string {
	if t.Branch == "" {
		return ""
	}

	return fmt.Sprintf("@branch:%s", t.Branch)
}

func (t Task) GetFormattedRepo() string {
	if t.Repo == "" {
		return ""
	}

	return fmt.Sprintf("@repo:%s", t.Repo)
}
