package internal

import (
	"fmt"
	"slices"
)

type TaskGroup struct {
	dateStr string
	tasks   []Task
}

func (g *TaskGroup) MatchDateStr(dateStr string) bool {
	if dateStr == g.dateStr {
		return true
	}

	return false
}

func RenderAsList(tasks []Task) {
	var groups []TaskGroup

	for _, task := range tasks {
		dateStr := task.GetDate()
		idx := slices.IndexFunc(groups, func(g TaskGroup) bool {
			return g.MatchDateStr(dateStr)
		})

		if idx != -1 {
			group := groups[idx]
			group.tasks = append(group.tasks, task)
			groups[idx] = group
		} else {
			group := TaskGroup{
				dateStr: dateStr,
				tasks:   []Task{task},
			}

			groups = append(groups, group)
		}
	}

	for _, group := range groups {
		fmt.Printf("%s\n", group.dateStr)

		for _, task := range group.tasks {
			fmt.Printf("    %s", task.ToStringWithHours())
		}
	}
}
