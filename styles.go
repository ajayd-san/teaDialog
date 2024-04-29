package main

import "github.com/charmbracelet/lipgloss"

var (
	containerStyle            = lipgloss.NewStyle().Align(lipgloss.Center).AlignVertical(lipgloss.Center)
	dialogStyle               = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("69")).Align(lipgloss.Center).Padding(3)
	promptStyle               = lipgloss.NewStyle().Align(lipgloss.Center)
	selectedPromptStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderRight(false).BorderTop(false).BorderBottom(false)
	selectedPromptOptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("98"))
)

const checkMark = "✔"
const checkBox = "▢"
