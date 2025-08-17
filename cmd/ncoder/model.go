package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	ind int
}

func (m model) Init() tea.Cmd {
	return nil
}
