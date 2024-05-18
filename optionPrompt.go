package teadialog

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type OptionPrompt struct {
	label          string
	options        []string
	selectedOption int
	cursorIndex    int
	isActive       bool
}

func (m OptionPrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, promptKeymap.Prev):
			if m.cursorIndex == 0 {
				m.cursorIndex = len(m.options) - 1
			} else {
				m.cursorIndex -= 1
			}

		case key.Matches(msg, promptKeymap.Next):
			if m.cursorIndex == len(m.options)-1 {
				m.cursorIndex = 0
			} else {
				m.cursorIndex += 1
			}
		case key.Matches(msg, promptKeymap.Select):
			m.selectedOption = m.cursorIndex
		}

	}
	return m, nil
}

func (m OptionPrompt) View() string {
	var res strings.Builder

	res.WriteString(m.label + "\n")
	for i, option := range m.options {
		var checkbox string
		checkbox = checkBox

		if i == m.selectedOption {
			checkbox = checkMark
		}

		if m.isActive && i == m.cursorIndex {
			str := fmt.Sprintf("%s %s\t", checkbox, option)
			res.WriteString(selectedPromptOptionStyle.Render(str))
		} else {
			res.WriteString(fmt.Sprintf("%s %s\t", checkbox, option))
		}
	}

	res.WriteString("\n\n")

	return promptStyle.Render(res.String())
}

func (m OptionPrompt) Init() tea.Cmd {
	return nil
}

//util

func (p OptionPrompt) setIsActive(state bool) Prompt {
	p.isActive = state
	return p
}

// func (p OptionPrompt) getIsActive() *bool {
// 	return &p.isActive
// }

func MakeOptionPrompt(message string, options []string) Prompt {
	return &OptionPrompt{
		label:          message,
		options:        options,
		selectedOption: -1,
	}
}
