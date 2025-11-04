package task

import (
	"testing"
	"time"
)

// 0001: [Sun, 31 Dec 1899 00:00:00 GMT] "XXX-123 this is my task" @tags:feat @issue_id:XXX-123 @repo:<name> @branch:feat/XXX-123-this-is-my-task
func TestParse(t *testing.T) {
	tests := map[string]struct {
		in  string
		out Task
	}{
		"basic task without tags, repo, branch": {
			in:  `001: [2006-01-02 15:04:00] ABC-123 "this is my task"`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC)},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			task, _ := StrToTask(test.in)

			AssetEqual(t, task.ID, test.out.ID)
			AssetEqual(t, task.Created, test.out.Created)
			AssetEqual(t, task.Branch, test.out.Branch)
			AssetEqual(t, task.Repo, test.out.Repo)
			AssetEqual(t, task.IssueId, test.out.IssueId)
		})
	}
}

func AssetEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}
