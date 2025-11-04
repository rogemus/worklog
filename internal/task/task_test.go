package task

import (
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	tests := map[string]struct {
		in  Task
		out string
	}{
		"basic task without tags, repo, branch": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name"},
			out: `001: [2006-01-02 15:04:00] Task Name`,
		},
		"task with tags": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Tags: []string{"feat", "ui"}},
			out: `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui`,
		},
		"task with branch": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master"},
			out: `001: [2006-01-02 15:04:00] Task Name @branch:master`,
		},
		"task with repo": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo"},
			out: `001: [2006-01-02 15:04:00] Task Name @repo:super-repo @branch:master`,
		},
		"task with branch and tags": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Tags: []string{"feat", "ui"}},
			out: `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @branch:master`,
		},
		"task with branch, repo, tags": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
			out: `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @repo:super-repo @branch:master`,
		},
		"task with branch, repo, issue_id, tags": {
			in:  Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), IssueId: "XXX-123", Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
			out: `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @issue_id:XXX-123 @repo:super-repo @branch:master`,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			taskStr := test.in.GetFormatted()

			AssetEqual(t, taskStr, test.out)
		})
	}
}
