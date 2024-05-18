package teadialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ErrorDialog struct {
	message string
}

const ErrorContinue = 0

func (e ErrorDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, NavKeymap.Enter):
			return nil, func() tea.Msg { return ErrorContinue }
		}

	}
	return e, nil
}

func (e ErrorDialog) View() string {

	var res strings.Builder

	res.WriteString("Error:\n\n")
	res.WriteString(e.message)

	return errorDialogStyle.Render(res.String())

}

func (e ErrorDialog) Init() tea.Cmd {
	return nil
}

// util
func NewErrorDialog(errMsg string) ErrorDialog {
	return ErrorDialog{message: errMsg}
}
