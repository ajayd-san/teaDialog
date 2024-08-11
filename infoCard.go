package teadialog

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Opt func(*InfoCard)

type InfoCard struct {
	title     string
	Message   string
	MinWidth  int
	MinHeight int
	storage   map[string]string
	help      help.Model
}

// Update implements tea.Model.
func (m InfoCard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, NavKeymap.Back, NavKeymap.Enter):
			return m, func() tea.Msg { return CloseDialog{} }
		}
	}

	return m, nil
}

func (m InfoCard) View() string {

	headerStyled := infoHeaderStyle.Render(m.title) + "\n"
	dialogBody := wordwrap.String(m.Message, m.MinWidth)

	infoStyleC := infoBodyStyle.Copy()
	if m.MinWidth != 0 {
		infoStyleC = infoStyleC.Width(m.MinWidth)
	}
	if m.MinHeight != 0 {
		infoStyleC = infoStyleC.Height(m.MinHeight)
	}

	styledDialogBody := infoStyleC.Render(dialogBody)
	headerWithBody := lipgloss.JoinVertical(lipgloss.Center, headerStyled, styledDialogBody)
	resFinal := infoStyle.Render(headerWithBody)
	dialogWithHelp := lipgloss.JoinVertical(lipgloss.Top, resFinal, "\n", m.help.View(infoCardKeymap))

	return dialogWithHelp
}

func (m InfoCard) Init() tea.Cmd {
	return nil
}

func InitInfoCard(title string, message string, kind DialogType, opts ...Opt) InfoCard {
	base := &InfoCard{
		title:   title,
		Message: message,
		help:    help.New(),
	}

	for _, opt := range opts {
		opt(base)
	}

	return *base
}

func (d InfoCard) GetStorage() map[string]string {
	return d.storage
}

// opt

// Set minimum width of the InfoCard
func WithMinHeight(minHeight int) Opt {
	return func(card *InfoCard) {
		card.MinHeight = minHeight
	}
}

// Set minimum width of the InfoCard
func WithMinWidth(minWidth int) Opt {
	return func(card *InfoCard) {
		card.MinWidth = minWidth
	}
}

// sets info card storage
func WithStorage(storage map[string]string) Opt {
	return func(card *InfoCard) {
		card.storage = storage
	}
}
