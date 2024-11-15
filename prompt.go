package teadialog

import (
	tea "github.com/charmbracelet/bubbletea"
)

type PromptInit int

type Prompt interface {
	tea.Model
	SetIsActive(bool) Prompt
	GetId() string
	GetSelection() any
	IsFocused() bool
}

type Hijacker interface {
	Prompt
	Hijack()
	UnHijack()
	// custom help
	Help() helpKeys
}
