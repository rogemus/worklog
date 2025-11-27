package reports

import (
	"fmt"
	"slices"
	"time"
	"worklog/internal/task"
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
		fmt.Printf("\n%s (%s)\n", group.dateStr, group.date.Weekday())

		for _, task := range group.tasks {
			fmt.Printf("  %s\n", task.ToString())
		}
	}
}
