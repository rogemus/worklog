package internal

import "github.com/charmbracelet/lipgloss"

var blue = lipgloss.Color("12")
var magenta = lipgloss.Color("5")
var dim = lipgloss.Color("250")

var itemStyles = lipgloss.NewStyle().
	PaddingRight(1).
	PaddingLeft(1)

var listTitleStyles = lipgloss.NewStyle().
	Foreground(magenta).
	PaddingBottom(1).
	PaddingTop(1)

var listHeadingStyles = lipgloss.NewStyle().
	Foreground(blue)

var dimStyles = lipgloss.NewStyle().
	Foreground(dim)

var titleStyles = lipgloss.NewStyle().
	PaddingLeft(2)
