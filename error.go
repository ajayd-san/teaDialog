package teadialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/muesli/reflow/wordwrap"
)

type ErrorDialog struct {
	message string
	help    help.Model
}

func (e ErrorDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, nil
}

func (e ErrorDialog) View() string {

	var res strings.Builder

	res.WriteString("Error:\n\n")
	res.WriteString(e.message)
	errDialog := errorDialogStyle.Render(res.String())
	finalStr := lipgloss.JoinVertical(lipgloss.Top, errDialog, "\n", e.help.View(errorKeyMap))

	return finalStr

}

func (e ErrorDialog) Init() tea.Cmd {
	return nil
}

// util
func NewErrorDialog(errMsg string, width int) ErrorDialog {
	//30 seems to be a sensible default
	errMsg = wordwrap.String(errMsg, width-30)
	help := help.New()
	help.Width = width
	return ErrorDialog{message: errMsg, help: help}
}
