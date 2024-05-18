package teadialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// type of dialog, can be used to distinguish between different dialogs in main update loop
type DialogType int

type Dialog struct {
	title        string
	prompts      []Prompt
	activePrompt int
	Kind         DialogType
	storage      map[string]string
}

type DialogSelectionResult struct {
	Kind        DialogType
	userChoices map[string]any
	userStorage map[string]string
}

// Update implements tea.Model.
func (m Dialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case PromptInit:
		//set the first prompt to active
		m.prompts[0] = m.prompts[0].setIsActive(true)

		/*
			find the maxwidth of each prompt rendered independently, this is ensures the dialog width is constant when left border
			is applied when a active prompt is displayed
		*/
		maxWidth := 0
		for _, prompt := range m.prompts {
			maxWidth = max(lipgloss.Width(dialogStyle.Render(selectedPromptStyle.Render(prompt.View()))), maxWidth)
		}

		dialogStyle = dialogStyle.Width(maxWidth)

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
		case key.Matches(msg, NavKeymap.Enter):
			return m, m.getUserChoices()
		}
	}

	updatedPrompt, _ := m.getActivePrompt().Update(msg)
	temp := updatedPrompt.(Prompt)
	m.prompts[m.activePrompt] = temp

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
	//INFO: fire init prompt command to initialize the prompt, duh
	return func() tea.Msg {
		return PromptInit(0)
	}
}

func InitDialogue(title string, prompts []Prompt, dialogKind DialogType, storage map[string]string) Dialog {
	return Dialog{
		title:        title,
		prompts:      prompts,
		activePrompt: 0,
		Kind:         dialogKind,
		storage:      storage,
	}
}

func (d Dialog) GetStorage() map[string]string {
	return d.storage
}

func (d *Dialog) SetStorage(storage map[string]string) {
	d.storage = storage
}

func (d *Dialog) getUserChoices() tea.Cmd {
	return func() tea.Msg {
		data := make(map[string]any, len(d.prompts))
		for _, prompt := range d.prompts {
			id := prompt.GetId()
			selection := prompt.GetSelection()
			data[id] = selection
		}

		return DialogSelectionResult{
			Kind:        d.Kind,
			userChoices: data,
			userStorage: d.storage,
		}
	}
}

// nav
func (d *Dialog) nextPrompt() {
	if !(d.activePrompt == len(d.prompts)-1) {
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].setIsActive(false)
		d.activePrompt += 1
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].setIsActive(true)
	}

}

func (d *Dialog) prevPrompt() {
	if !(d.activePrompt == 0) {
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].setIsActive(false)
		d.activePrompt -= 1
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].setIsActive(true)
	}
}

//util

func (d Dialog) getActivePrompt() Prompt {
	return d.prompts[d.activePrompt]
}
