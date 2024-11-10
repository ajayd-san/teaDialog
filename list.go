package teadialog

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var buttonStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#888B7E")).
	Padding(0, 3).
	Margin(1, 0)

var activeButtonStyle = buttonStyle.
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#F25D94")).
	Margin(1, 1, 1, 0)

type PopupListItem struct {
	Name, Desc string
}

func (i PopupListItem) Title() string       { return i.Name }
func (i PopupListItem) Description() string { return i.Desc }
func (i PopupListItem) FilterValue() string { return i.Name }

type PopupList struct {
	id            string
	list          list.Model
	buttonMsg     string
	isActive      bool
	selection     *PopupListItem
	isHijacker    bool
	isListFocused bool
}

func (m *PopupList) Hijack() {
	m.isHijacker = true
}

func (m *PopupList) UnHijack() {
	m.isHijacker = false
}

func (m *PopupList) SetIsActive(state bool) Prompt {
	m.isActive = state
	return m
}

func (m *PopupList) GetId() string {
	return m.id
}

func (m *PopupList) GetSelection() any {
	return m.selection.Name
}

func (m *PopupList) IsFocused() bool {
	return m.isListFocused
}

func (m *PopupList) Init() tea.Cmd {
	return nil
}

func (m *PopupList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, promptKeymap.Select):
			// if list is already hijacking the dialog then pressing space will select the item
			if m.isHijacker {
				if selected, ok := m.list.SelectedItem().(PopupListItem); ok {
					m.selection = &selected
					m.isListFocused = false
					return m, func() tea.Msg { return takeControlBackMsg{} }
				}
			} else {
				m.isListFocused = true
				return m, func() tea.Msg { return HijackMsg{} }
			}

		case key.Matches(msg, NavKeymap.Back) && m.isListFocused:
			m.isListFocused = false
			m.selection = nil
			return m, func() tea.Msg { return takeControlBackMsg{} }
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *PopupList) View() string {
	if m.isListFocused {
		return m.list.View()
	}

	if !m.isHijacker && m.selection != nil {
		return "Selected: " + selectedPromptOptionStyle.Render(m.selection.Name)
	} else if !m.isHijacker {
		if m.isActive {
			return activeButtonStyle.Render(m.buttonMsg)
		}
		return buttonStyle.Render(m.buttonMsg)
	}

	return ""
}

func Default_list(items []list.Item, button_msg string, width, height int) PopupList {
	del := list.NewDefaultDelegate()
	del.ShowDescription = false

	m := PopupList{
		list:      list.New(items, del, 40, 15),
		buttonMsg: "Select Pod",
	}
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)
	m.list.DisableQuitKeybindings()

	return m
}
