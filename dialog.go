package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Dialog struct {
	title   string
	prompts []Prompt
}

// Update implements tea.Model.
func (m Dialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width).Height(msg.Height)

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Dialog) View() string {

	var res strings.Builder

	res.WriteString(m.title + "\n\n")

	for _, prompt := range m.prompts {
		res.WriteString(prompt.View())
	}

	return containerStyle.Render(dialogStyle.Render(res.String()))
}

func (m Dialog) Init() tea.Cmd {
	return nil
}

func InitDialogue() Dialog {
	return Dialog{title: "test", prompts: []Prompt{MakePrompt("are your sure man, this is a hard choice?", []string{"yes", "no"})}}
}
