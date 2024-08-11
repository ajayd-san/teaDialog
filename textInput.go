package teadialog

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputOpt func(*TextInput)

type TextInput struct {
	id       string
	input    textinput.Model
	label    string
	isActive bool
	defaults string
}

func (ti TextInput) SetIsActive(state bool) Prompt {
	if state {
		ti.input.Focus()
	} else {
		ti.input.Blur()
	}
	ti.isActive = state
	return ti
}

func (ti TextInput) GetId() string {
	return ti.id
}

func (te TextInput) GetSelection() any {
	return te.input.Value()
}

func MakeTextInputPrompt(id string, label string, opts ...TextInputOpt) TextInput {
	ti := textinput.New()
	ti.CharLimit = 156
	ti.Width = 20

	m := TextInput{
		input: ti,
		id:    id,
		label: label,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m

}

func (m TextInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyEsc:
			m.input.Blur()
			return m, nextPrompt
		}
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m TextInput) IsFocused() bool {
	return m.input.Focused()
}

func (m TextInput) View() string {
	label := m.label
	if m.IsFocused() {
		label = selectedPromptOptionStyle.Render(m.label)
	}

	return promptStyle.Render(
		fmt.Sprintf(
			"%s\n\n%s\n\n",
			label,
			m.input.View(),
		),
	)

}

// Opts
func WithPlaceHolder(placeholder string) TextInputOpt {
	return func(textinput *TextInput) {
		textinput.input.Placeholder = placeholder
	}
}

func WithDefaultValue(defaults string) TextInputOpt {
	return func(textinput *TextInput) {
		textinput.defaults = defaults
	}
}
