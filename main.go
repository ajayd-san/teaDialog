package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(InitDialogue(), tea.WithAltScreen())
	f, _ := tea.LogToFile("debug.log", "debug")
	defer f.Close()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
