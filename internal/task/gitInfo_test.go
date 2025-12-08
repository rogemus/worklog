package task

import (
	"testing"

	"github.com/rogemus/worklog/internal/utils"
)

func TestGitToTask(t *testing.T) {
	tests := map[string]struct {
		in  []string
		out Task
	}{
		"feat branch with long task ID": {
			in:  []string{"feat/XXX-1-add-login", "auth-service"},
			out: Task{Branch: "feat/XXX-1-add-login", Name: "XXX-1: Add Login", IssueId: "XXX-1", Tags: []string{"feat"}, Repo: "auth-service"},
		},
		"fix branch with short task ID": {
			in:  []string{"fix/2-correct-validation", "user-service"},
			out: Task{Branch: "fix/2-correct-validation", Name: "2: Correct Validation", IssueId: "2", Tags: []string{"fix"}, Repo: "user-service"},
		},
		"FEAT branch uppercase long task ID": {
			in:  []string{"FEAT/XXX-3-ADD-PAYMENT-GATEWAY", "payment-service"},
			out: Task{Branch: "FEAT/XXX-3-ADD-PAYMENT-GATEWAY", Name: "XXX-3: Add Payment Gateway", IssueId: "XXX-3", Tags: []string{"feat"}, Repo: "payment-service"},
		},
		"FIX branch uppercase short task ID": {
			in:  []string{"FIX/4-HANDLE-TIMEOUT", "order-service"},
			out: Task{Branch: "FIX/4-HANDLE-TIMEOUT", Name: "4: Handle Timeout", IssueId: "4", Tags: []string{"fix"}, Repo: "order-service"},
		},
		"feat branch without task ID": {
			in:  []string{"feat/improve-dashboard", "frontend-app"},
			out: Task{Branch: "feat/improve-dashboard", Name: "Improve Dashboard", Tags: []string{"feat"}, Repo: "frontend-app"},
		},
		"FEAT branch without task ID": {
			in:  []string{"FEAT/UPDATE-API-ENDPOINTS", "api-gateway"},
			out: Task{Branch: "FEAT/UPDATE-API-ENDPOINTS", Name: "Update Api Endpoints", Tags: []string{"feat"}, Repo: "api-gateway"},
		},
		"simple branch lowercase without task ID": {
			in:  []string{"add-login", "auth-service"},
			out: Task{Branch: "add-login", Name: "Add Login", Repo: "auth-service"},
		},
		"simple branch uppercase without task ID": {
			in:  []string{"ADD-PAYMENT-GATEWAY", "payment-service"},
			out: Task{Branch: "ADD-PAYMENT-GATEWAY", Name: "Add Payment Gateway", Repo: "payment-service"},
		},
		"simple branch lowercase with short task ID": {
			in:  []string{"123-add-login", "auth-service"},
			out: Task{Branch: "123-add-login", Name: "123: Add Login", IssueId: "123", Repo: "auth-service"},
		},
		"simple branch lowercase with long task ID": {
			in:  []string{"ABC-123-add-login", "auth-service"},
			out: Task{Branch: "ABC-123-add-login", Name: "ABC-123: Add Login", IssueId: "ABC-123", Repo: "auth-service"},
		},
		"simple branch uppercase with short task ID": {
			in:  []string{"123-ADD-PAYMENT-GATEWAY", "payment-service"},
			out: Task{Branch: "123-ADD-PAYMENT-GATEWAY", Name: "123: Add Payment Gateway", IssueId: "123", Repo: "payment-service"},
		},
		"simple branch uppercase with long task ID": {
			in:  []string{"CDA-123-ADD-PAYMENT-GATEWAY", "payment-service"},
			out: Task{Branch: "CDA-123-ADD-PAYMENT-GATEWAY", Name: "CDA-123: Add Payment Gateway", IssueId: "CDA-123", Repo: "payment-service"},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			task := GitToTask(test.in[0], test.in[1])
			utils.AssetEqual(t, task.ID, 0)
			utils.AssetEqual(t, task.Created, test.out.Created)
			utils.AssetEqual(t, task.Name, test.out.Name)
			utils.AssetEqual(t, task.Branch, test.out.Branch)
			utils.AssetEqual(t, task.Repo, test.out.Repo)
			utils.AssetEqual(t, task.IssueId, test.out.IssueId)
			utils.AssetEqualSlice(t, task.Tags, test.out.Tags)
		})
	}
}
