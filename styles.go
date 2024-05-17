package teadialog

import "github.com/charmbracelet/lipgloss"

var (
	containerStyle            = lipgloss.NewStyle().Align(lipgloss.Center).AlignVertical(lipgloss.Center)
	dialogStyle               = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("69")).Align(lipgloss.Center).Padding(2, 8)
	promptContainerStyle      = lipgloss.NewStyle().Align(lipgloss.Left)
	promptStyle               = lipgloss.NewStyle().MarginLeft(1)
	selectedPromptStyle       = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).Padding(0, 0, 0, 1)
	selectedPromptOptionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("98"))
)

const checkMark = "✔"
const checkBox = "▢"
