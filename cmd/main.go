package main

import (
	"AdventOfCode/pkg/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	years := tui.NewYearsModel()
	program := tea.NewProgram(years)
	_, err := program.Run()

	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
