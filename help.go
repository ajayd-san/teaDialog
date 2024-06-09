package teadialog

import (
	"github.com/charmbracelet/bubbles/key"
)

type helpKeys struct {
	Navigation key.Binding
	Back       key.Binding
	Submit     key.Binding
	Select     key.Binding
}

func (k helpKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Navigation, k.Select, k.Submit, k.Back}
}

func (k helpKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

var helpKeyMap = helpKeys{
	Navigation: key.NewBinding(
		key.WithKeys("h/j/k/l/arrowKeys"),
		key.WithHelp("Arrow Keys/h/j/k/l", "Navigation"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
	),
	Submit: key.NewBinding(
		key.WithKeys("Enter"),
		key.WithHelp("Enter", "Submit"),
	),
	Select: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("Space", "Select Option"),
	),
}
