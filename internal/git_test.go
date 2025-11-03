package internal

import "testing"

func TestParseBranchName(t *testing.T) {

	var tests = map[string]struct {
		in      string
		outName string
		outTags []Tag
	}{
		"with type and issue id": {
			in:      "feat/XXX-123-feature",
			outName: "featue",
			outTags: []Tag{Tag("feat"), Tag("XXX-123")},
		},
		"with invalid type and valid issue id": {
			in:      "feat2/XXX-123-feature-with-long",
			outName: "feature with long",
			outTags: []Tag{Tag("XXX-123")},
		},
		"with only issue id": {
			in:      "XXX-123-feature-with-long-name",
			outName: "featue with long name",
			outTags: []Tag{Tag("XXX-123")},
		},
		"with only invalid issue id": {
			in:      "XXX123-feature-with-long-name",
			outName: "XXX123 featue with long name",
			outTags: []Tag{},
		},
		"replace `-` with space": {
			in:      "super-long-branch-name-with-spaces",
			outName: "super long branch name with spaces",
			outTags: []Tag{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			taskName, taskTags := ParseBranchName(test.in)
			AssertEqualSlice(t, taskTags, test.outTags)
			AssertEqual(t, taskName, test.outName)
		})
	}
}
