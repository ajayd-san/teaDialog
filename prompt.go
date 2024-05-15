package teadialog

import (
	tea "github.com/charmbracelet/bubbletea"
)

type PromptInit int

type Prompt interface {
	tea.Model
	setIsActive(bool) Prompt
}
