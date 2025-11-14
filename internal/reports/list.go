package reports

import (
	"fmt"
	"slices"
	"worklog/internal/task"
)

type TaskGroup struct {
	dateStr string
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
				tasks:   []task.Task{t},
			}

			groups = append(groups, group)
		}
	}

	for _, group := range groups {
		fmt.Printf("%s\n", group.dateStr)

		for _, task := range group.tasks {
			fmt.Printf("   [%s] %s\n", task.GetFormattedCreatedTime(), task.Name)
		}
	}
}
