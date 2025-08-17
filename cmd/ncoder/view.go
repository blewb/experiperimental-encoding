package main

import "strings"

func (m model) View() string {

	s := "\n"

	edgeRow := borderStyle.Render(" +" + strings.Repeat("-", INNER_WIDTH) + "+")
	emptyRow := borderStyle.Render(" |" + strings.Repeat(" ", INNER_WIDTH) + "|")

	s += edgeRow + "\n"

	s += emptyRow + "\n"
	s += emptyRow + "\n"
	s += emptyRow + "\n"
	s += emptyRow + "\n"

	s += edgeRow + "\n"

	s += "\n\n"

	return s

}
