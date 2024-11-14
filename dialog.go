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

type DialogOption func(*Dialog)

type nextprompt struct{}

type CloseDialog struct{}

type QuitDialog struct{}

type HijackMsg struct{}

type takeControlBackMsg struct{}

type Dialog struct {
	title   string
	prompts []Prompt
	// used to indicate if a nested dialog is in use, so that we can use full space
	hijacked     bool
	hijacker     Hijacker
	activePrompt int
	Kind         DialogType
	storage      map[string]string
	help         help.Model
	helpKeymap   helpKeys
	width        int
	height       int
}

type DialogSelectionResult struct {
	Kind        DialogType
	UserChoices map[string]any
	UserStorage map[string]string
}

// Update implements tea.Model.
func (m Dialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {

	case PromptInit:
		//set the first prompt to active

		if len(m.prompts) != 0 {
			m.prompts[0] = m.prompts[0].SetIsActive(true)
			cmds = append(cmds, m.prompts[0].Init())
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
		if m.width == 0 {
			tempdialogStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("69")).Align(lipgloss.Center).Padding(2, 8)
			for _, prompt := range m.prompts {
				maxWidth = max(lipgloss.Width(tempdialogStyle.Render(selectedPromptStyle.Render(prompt.View()))), maxWidth)
			}
		} else {
			maxWidth = m.width
		}

		promptContainerStyle = promptContainerStyle.Width(maxWidth)

		if m.height != 0 {
			dialogStyle = dialogStyle.Height(m.height)
		}

	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width).Height(msg.Height)
		m.help.Width = msg.Width

	case nextprompt:
		if m.activePrompt == len(m.prompts)-1 {
			return m, func() tea.Msg { return CloseDialog{} }
		}

		m.nextPrompt()
		cmd := m.getActivePrompt().Init()
		cmds = append(cmds, cmd)

	case HijackMsg:
		if hijacker, ok := m.prompts[m.activePrompt].(Hijacker); ok {
			m.hijacked = true
			m.hijacker = hijacker
			m.hijacker.Hijack()
		}
	case takeControlBackMsg:
		m.hijacked = false
		m.hijacker.UnHijack()
		m.hijacker = nil

	case tea.KeyMsg:

		if !m.hijacked {
			switch {
			case len(m.prompts) > 0 && key.Matches(msg, NavKeymap.Next) && !m.getActivePrompt().IsFocused():
				m.nextPrompt()
				cmd := m.getActivePrompt().Init()
				cmds = append(cmds, cmd)
				return m, nil
			case len(m.prompts) > 0 && key.Matches(msg, NavKeymap.Prev) && !m.getActivePrompt().IsFocused():
				m.prevPrompt()
				cmd := m.getActivePrompt().Init()
				cmds = append(cmds, cmd)
				return m, nil
			case key.Matches(msg, NavKeymap.Enter):
				if !m.getActivePrompt().IsFocused() {
					return m, func() tea.Msg { return CloseDialog{} }
				}
			case key.Matches(msg, NavKeymap.SkipAndSubmit):
				return m, func() tea.Msg { return CloseDialog{} }
			case key.Matches(msg, NavKeymap.Back):
				return m, func() tea.Msg { return QuitDialog{} }
			}
		}
	}

	if m.hijacked {
		tempm, cmd := m.hijacker.Update(msg)
		cmds = append(cmds, cmd)
		if h, ok := tempm.(Hijacker); ok {
			m.hijacker = h
		}
	} else {
		if len(m.prompts) != 0 {
			updatedPrompt, cmd := m.getActivePrompt().Update(msg)
			temp := updatedPrompt.(Prompt)
			m.prompts[m.activePrompt] = temp
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Dialog) View() string {

	var res strings.Builder
	var promptStrs strings.Builder

	if m.hijacked {
		v := m.hijacker.View()
		res.WriteString(v)
	} else {
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
	}

	dialogFinal := dialogStyle.Render(res.String())
	dialogWithHelp := lipgloss.JoinVertical(lipgloss.Center, dialogFinal, "\n", m.help.View(m.helpKeymap))

	return containerStyle.Render(dialogWithHelp)
}

func (m Dialog) Init() tea.Cmd {
	//INFO: fire init prompt command to initialize the prompt, duh
	return func() tea.Msg {
		return PromptInit(0)
	}
}

func InitDialogWithPrompt(title string, prompts []Prompt, dialogKind DialogType, storage map[string]string, opts ...DialogOption) Dialog {
	m := Dialog{
		title:        title,
		prompts:      prompts,
		activePrompt: 0,
		Kind:         dialogKind,
		storage:      storage,
		help:         help.New(),
		helpKeymap:   helpKeyMap,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}

func WithShowFullHelp(show bool) DialogOption {
	return func(d *Dialog) {
		d.helpKeymap.showFull = show
	}
}
func WithWidth(width int) DialogOption {
	return func(d *Dialog) {
		d.width = width
	}
}

func WithHeight(height int) DialogOption {
	return func(d *Dialog) {
		d.height = height
	}
}

func (d Dialog) GetStorage() map[string]string {
	return d.storage
}

func (d *Dialog) SetStorage(storage map[string]string) {
	d.storage = storage
}

func (d Dialog) GetUserChoices() DialogSelectionResult {
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

// nav
func (d *Dialog) nextPrompt() {
	if !(d.activePrompt == len(d.prompts)-1) {
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].SetIsActive(false)
		d.activePrompt += 1
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].SetIsActive(true)
	}

}

func (d *Dialog) prevPrompt() {
	if !(d.activePrompt == 0) {
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].SetIsActive(false)
		d.activePrompt -= 1
		d.prompts[d.activePrompt] = d.prompts[d.activePrompt].SetIsActive(true)
	}
}

//util

func (d Dialog) getActivePrompt() Prompt {
	return d.prompts[d.activePrompt]
}

func nextPrompt() tea.Msg {
	return nextprompt{}
}
