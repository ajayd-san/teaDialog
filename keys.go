package teadialog

import "github.com/charmbracelet/bubbles/key"

type navigationKeymap struct {
	Back          key.Binding
	Enter         key.Binding
	Quit          key.Binding
	Next          key.Binding
	Prev          key.Binding
	SkipAndSubmit key.Binding
}

type promptNavKeymap struct {
	Select key.Binding
	Next   key.Binding
	Prev   key.Binding
}

var promptKeymap = promptNavKeymap{
	Select: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "select"),
	),
	Next: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("->/l", "next option"),
	),
	Prev: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("<-/h", "prev option"),
	),
}

var NavKeymap = navigationKeymap{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
	Next: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("->/j", "next prompt"),
	),
	Prev: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("<-/k", "prev prompt"),
	),
	SkipAndSubmit: key.NewBinding(
		key.WithKeys("ctrl+a"),
		key.WithHelp("C-a", "Skip and submit"),
	),
}
