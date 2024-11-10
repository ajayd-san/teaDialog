package teadialog

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	id            string
	list          list.Model
	isActive      bool
	selection     *item
	isHijacker    bool
	isListFocused bool
}

func (m *model) Hijack() {
	m.isHijacker = true
}

func (m *model) UnHijack() {
	m.isHijacker = false
}

func (m *model) SetIsActive(state bool) Prompt {
	m.isActive = state
	return m
}

func (m *model) GetId() string {
	return m.id
}

func (m *model) GetSelection() any {
	return m.selection.title
}

// TODO: WRITE CORRECT LOGIC
func (m *model) IsFocused() bool {
	return m.isListFocused
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, promptKeymap.Select):
			// if list is already hijacking the dialog then pressing space will select the item
			if m.isHijacker {
				if selected, ok := m.list.SelectedItem().(item); ok {
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

		}
		// case tea.WindowSizeMsg:
		// h, v := docStyle.GetFrameSize()
		// m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *model) View() string {
	if m.isListFocused {
		return m.list.View()
	}
	if !m.isHijacker && m.selection != nil {
		return m.selection.title
	} else if !m.isHijacker {
		return "select"
	}

	return ""
}

func Default_list() model {
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "idk", desc: "It's good on toast"},
	}

	del := list.NewDefaultDelegate()
	del.ShowDescription = false

	m := model{list: list.New(items, del, 40, 15)}
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)
	m.list.DisableQuitKeybindings()

	return m
}
