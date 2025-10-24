package internal

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func RenderAsTable(tasks []Task) string {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderTop(false).BorderBottom(false).BorderLeft(false).BorderRight(false).
		BorderStyle(lipgloss.NewStyle()).
		Headers("ID", "CREATED", "TASK NAME").
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case table.HeaderRow:
				return headerStyle
			default:
				return cellStyle
			}
		})

	for _, task := range tasks {
		t.Row(
			strconv.Itoa(int(task.Id)),
			fmt.Sprintf("%s %s", task.GetDate(), task.GetHour()),
			task.Name,
		)
	}

	return t.String()
}
