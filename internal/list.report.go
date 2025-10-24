package internal

import (
	"slices"

	"github.com/charmbracelet/lipgloss/list"
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

func RenderAsList(tasks []Task) string {
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

	l := list.New().
		Enumerator(func(items list.Items, index int) string { return "" })

	for _, group := range groups {
		tasksList := list.New().
			Enumerator(func(items list.Items, index int) string { return "" })

		for _, task := range group.tasks {
			tasksList.Item(task.RenderWithHours())
		}

		l.Items(listHeadingStyles.Render(group.dateStr), tasksList, "")
	}

	return l.String()
}
