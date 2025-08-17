package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	CHAR_COUNT  = 25
	INNER_WIDTH = 87
)

var (
	borderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#acacac"))
)

func main() {

	p := tea.NewProgram(model{
		ind: 0,
	})

	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}

}
