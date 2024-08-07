package teadialog

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type TogglePrompt struct {
	id       string
	label    string
	selected bool
	isActive bool
}

func (m TogglePrompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, promptKeymap.Select):
			m.selected = !m.selected
		}
	}
	return m, nil
}

func (m TogglePrompt) View() string {
	var res strings.Builder
	var checkbox string
	checkbox = checkBox

	if m.selected {
		checkbox = checkMark
	}

	str := fmt.Sprintf("%s %s", checkbox, m.label)
	if m.isActive {
		str = selectedPromptOptionStyle.Render(str)
	}
	res.WriteString(str)
	res.WriteString("\n\n")

	return promptStyle.Render(res.String())
}

func (m TogglePrompt) Init() tea.Cmd {
	return nil
}

func (p TogglePrompt) SetIsActive(state bool) Prompt {
	p.isActive = state
	return p
}

// interface Prompt
func (m TogglePrompt) GetId() string {
	return m.id
}

func (m TogglePrompt) GetSelection() any {
	// return bool
	return m.selected
}

func (m TogglePrompt) IsFocused() bool {
	return false
}

// util
func MakeTogglePrompt(id string, message string) TogglePrompt {
	return TogglePrompt{
		id:    id,
		label: message,
	}
}
