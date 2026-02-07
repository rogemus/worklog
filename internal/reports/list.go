package reports

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/rogemus/worklog/internal/task"
	"github.com/rogemus/worklog/internal/utils"
)

type TaskGroup struct {
	dateStr string
	date    time.Time
	tasks   []task.Task
}

func (g *TaskGroup) MatchDateStr(dateStr string) bool {
	if dateStr == g.dateStr {
		return true
	}

	return false
}

func RenderAsList(tasks []task.Task) {
	var groups []TaskGroup

	for _, t := range tasks {
		dateStr := t.GetFormattedCreatedDate()
		idx := slices.IndexFunc(groups, func(g TaskGroup) bool {
			return g.MatchDateStr(dateStr)
		})

		if idx != -1 {
			group := groups[idx]
			group.tasks = append(group.tasks, t)
			groups[idx] = group
		} else {
			group := TaskGroup{
				dateStr: dateStr,
				date:    t.Created,
				tasks:   []task.Task{t},
			}

			groups = append(groups, group)
		}
	}

	for _, group := range groups {
		fmt.Printf("\n %s (%s) #%d ───────────────────────────────────────────────────── \n", group.dateStr, group.date.Weekday(), len(group.tasks))

		for _, t := range group.tasks {
			name := t.Name + strings.Repeat(" ", utils.MaxInt(40-len(t.Name), 0))
			id := t.IssueId + strings.Repeat(" ", utils.MaxInt(10-len(t.IssueId), 0))
			repo := t.Repo + strings.Repeat(" ", utils.MaxInt(15-len(t.Repo), 0))
			tags := strings.Join(t.Tags, ",")

			fmt.Printf("   %s %s %s [%s]\n", id, name, repo, tags)
		}
	}
}
