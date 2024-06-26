package teadialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
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
	help         help.Model
}

type DialogSelectionResult struct {
	Kind        DialogType
	UserChoices map[string]any
	UserStorage map[string]string
}

// Update implements tea.Model.
func (m Dialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case PromptInit:
		//set the first prompt to active

		if len(m.prompts) != 0 {
			m.prompts[0] = m.prompts[0].setIsActive(true)
		}

		/*
			find the maxwidth of each prompt rendered independently, this is ensures the dialog width is constant when left border
			is applied when a active prompt is displayed
		*/
		maxWidth := 0

		/*
			use the default style of `dialogStyle` when deciding the width of the new dialog, otherwise previous prompt dialogstyle is applied
			when a new prompt is called, otherwise new prompts are wrongly rendered
		*/
		tempdialogStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("69")).Align(lipgloss.Center).Padding(2, 8)
		for _, prompt := range m.prompts {
			maxWidth = max(lipgloss.Width(tempdialogStyle.Render(selectedPromptStyle.Render(prompt.View()))), maxWidth)
		}

		dialogStyle = dialogStyle.Width(maxWidth)

	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width).Height(msg.Height)
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case len(m.prompts) > 0 && key.Matches(msg, NavKeymap.Next):
			m.nextPrompt()
			return m, nil
		case len(m.prompts) > 0 && key.Matches(msg, NavKeymap.Prev):
			m.prevPrompt()
			return m, nil
		case key.Matches(msg, NavKeymap.Enter):
			return m, m.getUserChoices()
		}
	}

	if len(m.prompts) != 0 {
		updatedPrompt, _ := m.getActivePrompt().Update(msg)
		temp := updatedPrompt.(Prompt)
		m.prompts[m.activePrompt] = temp
	}

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
	dialogFinal := dialogStyle.Render(res.String())
	dialogWithHelp := lipgloss.JoinVertical(lipgloss.Center, dialogFinal, "\n", m.help.View(helpKeyMap))

	return containerStyle.Render(dialogWithHelp)
}

func (m Dialog) Init() tea.Cmd {
	//INFO: fire init prompt command to initialize the prompt, duh
	return func() tea.Msg {
		return PromptInit(0)
	}
}

func InitDialogWithPrompt(title string, prompts []Prompt, dialogKind DialogType, storage map[string]string) Dialog {
	return Dialog{
		title:        title,
		prompts:      prompts,
		activePrompt: 0,
		Kind:         dialogKind,
		storage:      storage,
		help:         help.New(),
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
			UserChoices: data,
			UserStorage: d.storage,
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
