package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Dialog struct {
	title        string
	prompts      []Prompt
	activePrompt int
}

// Update implements tea.Model.
func (m Dialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width).Height(msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, NavKeymap.Quit):
			return m, tea.Quit

		case key.Matches(msg, NavKeymap.Next):
			m.nextPrompt()
		case key.Matches(msg, NavKeymap.Back):
			m.prevPrompt()
		}
	}

	updatedPrompt, _ := m.getActivePrompt().Update(msg)
	m.prompts[m.activePrompt] = updatedPrompt.(Prompt)

	return m, nil
}

func (m Dialog) View() string {

	var res strings.Builder

	res.WriteString(m.title + "\n\n")

	for _, prompt := range m.prompts {
		res.WriteString(prompt.View())
		res.WriteString("\n")
	}

	return containerStyle.Render(dialogStyle.Render(res.String()))
}

func (m Dialog) Init() tea.Cmd {
	return nil
}

func InitDialogue() Dialog {
	return Dialog{title: "test", prompts: []Prompt{
		MakePrompt("are your sure man, this is a hard choice?", []string{"yes", "no"}),
		MakePrompt("are your sure man, this is a most defo hard choice?", []string{"yasss", "naah"}),
	}}
}

// nav
func (d *Dialog) nextPrompt() {
	if !(d.activePrompt == len(d.prompts)-1) {
		d.activePrompt += 1
	}
}

func (d *Dialog) prevPrompt() {
	if !(d.activePrompt == 0) {
		d.activePrompt -= 1
	}
}

//util

func (d Dialog) getActivePrompt() *Prompt {
	return &d.prompts[d.activePrompt]
}
