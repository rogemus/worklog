package task

import (
	"testing"
	"time"
	"worklog/internal/utils"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		in  string
		out Task
	}{
		"basic task without tags, repo, branch": {
			in:  `[2006-01-02 15:04:00] Task Name`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name"},
		},
		"task with tags": {
			in:  `[2006-01-02 15:04:00] Task Name @tags:feat,ui`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Tags: []string{"feat", "ui"}},
		},
		"task with branch": {
			in:  `[2006-01-02 15:04:00] Task Name @branch:master`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master"},
		},
		"task with repo": {
			in:  `[2006-01-02 15:04:00] Task Name @repo:super-repo @branch:master`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo"},
		},
		"task with branch and tags": {
			in:  `[2006-01-02 15:04:00] Task Name @tags:feat,ui @branch:master`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Tags: []string{"feat", "ui"}},
		},
		"task with branch, repo, tags": {
			in:  `[2006-01-02 15:04:00] Task Name @tags:feat,ui @repo:super-repo @branch:master`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
		},
		"task with branch, repo, issue_id, tags": {
			in:  `[2006-01-02 15:04:00] Task Name @tags:feat,ui @issue_id:XXX-123 @repo:super-repo @branch:master`,
			out: Task{Created: time.Date(2006, 1, 2, 15, 4, 0, 0, time.UTC), IssueId: "XXX-123", Name: "Task Name", Branch: "master", Repo: "super-repo", Tags: []string{"feat", "ui"}},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			task, _ := StrToTask(test.in, 0)

			utils.AssetEqual(t, task.ID, 0)
			utils.AssetEqual(t, task.Created, test.out.Created)
			utils.AssetEqual(t, task.Branch, test.out.Branch)
			utils.AssetEqual(t, task.Repo, test.out.Repo)
			utils.AssetEqual(t, task.IssueId, test.out.IssueId)
			utils.AssetEqualSlice(t, task.Tags, test.out.Tags)
		})
	}
}
