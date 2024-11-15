package teadialog

import (
	"github.com/charmbracelet/bubbles/key"
)

type helpKeys struct {
	showFull   bool
	Navigation key.Binding
	Back       key.Binding
	Submit     key.Binding
	// like submit but accepts defaults and submits
	SkipAndSubmit key.Binding
	Select        key.Binding
}

func (k helpKeys) ShortHelp() []key.Binding {
	if k.showFull {
		return []key.Binding{k.Navigation, k.Select, k.Submit, k.SkipAndSubmit, k.Back}
	}
	return []key.Binding{k.Navigation, k.Select, k.Submit, k.Back}
}

func (k helpKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

// var helpKeyMap = helpKeys{
// 	showFull: false,
// 	Navigation: key.NewBinding(
// 		key.WithKeys("h/j/k/l"),
// 		key.WithHelp("Arrow Keys/h/j/k/l", "Navigation"),
// 	),
// 	Back: key.NewBinding(
// 		key.WithKeys("esc"),
// 		key.WithHelp("esc", "Back"),
// 	),
// 	Submit: key.NewBinding(
// 		key.WithKeys("Enter"),
// 		key.WithHelp("Enter", "Submit"),
// 	),
// 	SkipAndSubmit: key.NewBinding(
// 		key.WithKeys("ctrl+a"),
// 		key.WithHelp("C-a", "Skip and submit"),
// 	),
// 	Select: key.NewBinding(
// 		key.WithKeys(" "),
// 		key.WithHelp("Space", "Select Option"),
// 	),
// }

func default_Help() helpKeys {
	return helpKeys{
		showFull: false,
		Navigation: key.NewBinding(
			key.WithKeys("h/j/k/l"),
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
		SkipAndSubmit: key.NewBinding(
			key.WithKeys("ctrl+a"),
			key.WithHelp("C-a", "Skip and submit"),
		),
		Select: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("Space", "Select Option"),
		),
	}
}

type errorKeys struct {
	Back key.Binding
}

var errorKeyMap = errorKeys{
	Back: key.NewBinding(
		key.WithKeys("esc", "enter"),
		key.WithHelp("Enter/esc", "Back"),
	),
}

func (k errorKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Back}
}

func (k errorKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

type infoCardKeys struct {
	Back key.Binding
}

var infoCardKeymap = infoCardKeys{
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("Enter/esc", "Back"),
	),
}

func (k infoCardKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Back}
}

func (k infoCardKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}
