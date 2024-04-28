package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Prompt struct {
	message  string
	options  []string
	selected int
	cursor   int
}

func (m Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch msg := msg.(type) {
	// }
	return m, nil
}

func (m Prompt) View() string {
	var res strings.Builder

	res.WriteString(m.message + "\n")
	for _, option := range m.options {
		res.WriteString(fmt.Sprintf("%s %s\t", checkBox, option))
	}

	return promptStyle.Render(res.String())
}

func (m Prompt) Init() tea.Cmd {
	return nil
}

func MakePrompt(message string, options []string) Prompt {
	return Prompt{
		message: message,
		options: options,
	}
}
