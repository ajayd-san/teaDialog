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
			return m, nil
		case key.Matches(msg, NavKeymap.Prev):
			m.prevPrompt()
			return m, nil
		}
	}

	updatedPrompt, _ := m.getActivePrompt().Update(msg)
	m.prompts[m.activePrompt] = updatedPrompt.(Prompt)

	return m, nil
}

func (m Dialog) View() string {

	var res strings.Builder
	var promptStrs strings.Builder

	res.WriteString(m.title + "\n\n")

	for i, prompt := range m.prompts {
		promptStr := prompt.View()
		if i == m.activePrompt {
			promptStr = selectedPromptStyle.Render(promptStr)
		}
		promptStrs.WriteString(promptStr)
		promptStrs.WriteString("\n")
	}

	promptStrsFinal := promptContainerStyle.Render(promptStrs.String())
	res.WriteString(promptStrsFinal)

	return containerStyle.Render(dialogStyle.Render(res.String()))
}

func (m Dialog) Init() tea.Cmd {
	//set first prompt as active, to display the selection highlight
	m.prompts[0].setIsActive(true)

	//INFO: weird formatting issue with this code
	// maxBorderWidth := 0

	// for _, prompt := range m.prompts {
	// 	maxBorderWidth = max(maxBorderWidth, lipgloss.Width(prompt.View()))
	// }

	// log.Println(maxBorderWidth)

	// selectedPromptOptionStyle = selectedPromptOptionStyle.Width(maxBorderWidth)

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
		d.prompts[d.activePrompt].setIsActive(false)
		d.activePrompt += 1
		d.prompts[d.activePrompt].setIsActive(true)
	}

}

func (d *Dialog) prevPrompt() {
	if !(d.activePrompt == 0) {
		d.prompts[d.activePrompt].setIsActive(false)
		d.activePrompt -= 1
		d.prompts[d.activePrompt].setIsActive(true)
	}
}

//util

func (d Dialog) getActivePrompt() *Prompt {
	return &d.prompts[d.activePrompt]
}
