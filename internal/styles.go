package internal

import "github.com/charmbracelet/lipgloss"

// COLORS
var (
	blue    = lipgloss.Color("12")
	magenta = lipgloss.Color("5")
	dim     = lipgloss.Color("250")
)

// SHARED
var (
	dimStyles = lipgloss.NewStyle().
		Foreground(dim)
)

// REPORT LIST
var (
	listItemStyles = lipgloss.NewStyle().
			PaddingRight(1).
			PaddingLeft(1)

	listTitleStyles = lipgloss.NewStyle().
			Foreground(magenta).
			PaddingBottom(1).
			PaddingTop(1)

	listHeadingStyles = lipgloss.NewStyle().
				Foreground(blue)

	titleStyles = lipgloss.NewStyle().
			PaddingLeft(2)
)

// REPORT TABLE
var (
	headerStyle = lipgloss.NewStyle().
			Foreground(blue).
			Bold(true).
			Align(lipgloss.Center)

	cellStyle = lipgloss.NewStyle().
			Padding(0, 1)
)
