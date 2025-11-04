package task

import (
	"testing"
	"time"
)

// 0001: [Sun, 31 Dec 1899 00:00:00 GMT] XXX-123 this is my task @tags:feat @issue_id:XXX-123 @repo:<name> @branch:feat/XXX-123-this-is-my-task
func TestParse(t *testing.T) {
	tests := map[string]struct {
		in  string
		out Task
	}{
		"basic task without tags, repo, branch": {
			in:  `001: [2006-01-02 15:04:00] Task Name`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name"},
		},
		"task with tags": {
			in:  `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Tags: []string{"feat", "ui"}},
		},
		"task with branch": {
			in:  `001: [2006-01-02 15:04:00] Task Name @branch:master`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master"},
		},
		"task with repo": {
			in:  `001: [2006-01-02 15:04:00] Task Name @repo:super-repo @branch:master`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo"},
		},
		"task with branch and tags": {
			in:  `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @branch:master`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Tags: []string{"feat", "ui"}},
		},
		"task with branch, repo, tags": {
			in:  `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @repo:super-repo @branch:master`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
		},
		"task with branch, repo, issue_id, tags": {
			in:  `001: [2006-01-02 15:04:00] Task Name @tags:feat,ui @issue_id:XXX-123 @repo:super-repo @branch:master`,
			out: Task{ID: 1, Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), IssueId: "XXX-123", Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
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
			AssetEqualSlice(t, task.Tags, test.out.Tags)
		})
	}
}
