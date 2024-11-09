package teadialog

import (
	"log"

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
	id        string
	list      list.Model
	isActive  bool
	selection item
}

func (m model) SetIsActive(state bool) Prompt {
	m.isActive = state
	return m
}

func (m model) GetId() string {
	return m.id
}

func (m model) GetSelection() any {
	return m.selection.title
}

// TODO: WRITE CORRECT LOGIC
func (m model) IsFocused() bool {
	return false
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("in list update")
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == " " {
			return m, func() tea.Msg { return HijackMsg{} }
		}
		// case tea.WindowSizeMsg:
		// h, v := docStyle.GetFrameSize()
		// m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	log.Println("list; ", msg)
	m.list, cmd = m.list.Update(msg)

	log.Println("after: ", m.list.Cursor())

	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func Default_list() model {
	items := []list.Item{
		item{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{title: "Nutella", desc: "It's good on toast"},
		item{title: "idk", desc: "It's good on toast"},
	}

	del := list.NewDefaultDelegate()
	del.ShowDescription = false

	del.Styles.NormalTitle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#1a1a1a", Dark: "#dddddd"}).Padding(0)
	del.Styles.SelectedTitle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, true).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"}).
		Foreground(lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"})
	m := model{list: list.New(items, del, 40, 15)}
	m.list.SetShowHelp(false)
	m.list.SetShowTitle(false)

	return m
}
