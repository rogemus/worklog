package internal

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var branchTypes []string = []string{
	"feat",
	"fix",
	"docs",
	"style",
	"refactor",
	"test",
	"chore",
}

// r [A-Z]+-[0-9]+-
// feat/XXX-123-feature -> featue, Tag{feat, XXX-123}
// XXX-123-feature -> feature, Tag{XXX-123}

// r [0-9]+-
// 123-feature --> feature, Tag{123}

// feature-new -> feature new, Tag{}

// super-long-branch-name-with-spaces -> super long branch name with spaces
func ParseBranchName(branch string, repo string) (string, []Tag) {
	var name string
	var tags []Tag

	tags = append(
		tags,
		Tag(fmt.Sprintf("@_branch:%s", branch)),
		Tag(fmt.Sprintf("@_repo:%s", repo)),
	)

	name = strings.ReplaceAll(branch, "-", " ")

	regBranch, _ := regexp.Compile("[a-z]+/.+")
	regIssue, _ := regexp.Compile("[A-Z]+-[0-9]+-")
	regIssueNumber, _ := regexp.Compile("[0-9]+-")

	if regBranch.MatchString(branch) {
		parts := strings.Split(branch, "/")

		if slices.Contains(branchTypes, parts[0]) {
			tag := Tag(parts[0])
			tags = append(tags, tag)
		}

		if 

	}

	// if strings.Contains(branch, "/") {
	// 	parts := strings.Split(branch, "/")
	//
	// 	if slices.Contains(branchTypes, parts[0]) {
	// 		tag := Tag(parts[0])
	// 		tags = append(tags, tag)
	// 	}
	// }

	return name, tags
}
