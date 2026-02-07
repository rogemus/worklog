package version

import (
	"sort"
	"strconv"
	"strings"
)

type SemVer struct {
	Original string
	Major    int
	Minor    int
	Patch    int
}

func NewSemVer(v string) SemVer {
	sv := SemVer{Original: strings.TrimSpace(v)}
	parts := strings.Split(v, ".")
	v = strings.TrimPrefix(v, "v")

	if len(parts) > 0 {
		sv.Major, _ = strconv.Atoi(parts[0])
	}
	if len(parts) > 1 {
		sv.Minor, _ = strconv.Atoi(parts[1])
	}
	if len(parts) > 2 {
		sv.Patch, _ = strconv.Atoi(parts[2])
	}

	return sv
}

func sortSemVer(semVers []SemVer) []SemVer {
	sort.Slice(semVers, func(i, j int) bool {
		if semVers[i].Major != semVers[j].Major {
			return semVers[i].Major < semVers[j].Major
		}

		if semVers[i].Minor != semVers[j].Minor {
			return semVers[i].Minor < semVers[j].Minor
		}

		return semVers[i].Patch < semVers[j].Patch
	})

	return semVers
}
