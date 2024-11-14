package teadialog

import (
	"fmt"
	"strings"

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
	Name           string
	AdditionalData any
}

func (i PopupListItem) Title() string       { return i.Name }
func (i PopupListItem) Description() string { return "" }
func (i PopupListItem) FilterValue() string { return i.Name }

type PopupList struct {
	id            string
	title         string
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
	return m.selection
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
		var res strings.Builder
		res.WriteString(fmt.Sprintf("Select %s: \n", m.title))
		res.WriteString(m.list.View())
		return res.String()
	}

	if !m.isHijacker && m.selection != nil {
		return fmt.Sprintf("Selected %s: %s\n", m.title, selectedPromptOptionStyle.Render(m.selection.Name))
	} else if !m.isHijacker {
		if m.isActive {
			return activeButtonStyle.Render(m.buttonMsg)
		}
		return buttonStyle.Render(m.buttonMsg)
	}

	return ""
}

func Default_list(id string, title string, items []PopupListItem, button_msg string, width, height int) PopupList {
	del := list.NewDefaultDelegate()
	del.ShowDescription = false

	finalList := make([]list.Item, 0, len(items))

	for _, item := range items {
		finalList = append(finalList, item)
	}

	m := PopupList{
		id:        id,
		title:     title,
		list:      list.New(finalList, del, width, height),
		buttonMsg: "Select Pod",
	}
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)
	m.list.DisableQuitKeybindings()

	return m
}
