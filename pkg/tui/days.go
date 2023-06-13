package tui

import (
	"AdventOfCode/pkg/config"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type DaysModel struct {
	list         list.Model
	choice       string
	selectedYear string
}

func newDaysModel(year string) DaysModel {
	model := DaysModel{selectedYear: year}
	items := model.PopulateItems()
	list := createList(items, "Select day:", config.Days, daysHelpOptions)

	model.list = list

	return model
}

func (model DaysModel) Init() tea.Cmd {
	return nil
}

func (model DaysModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		model.list.SetWidth(msg.Width)
		return model, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return model, tea.Quit

		case "enter":
			i, ok := model.list.SelectedItem().(item)

			if ok {
				model.choice = string(i)
			}

			return model, tea.Quit
		}
	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)

	return model, cmd
}

func (model DaysModel) View() string {
	return fmt.Sprintf("\n%s\n", model.list.View())
}

func (model DaysModel) PopulateItems() []list.Item {
	items := []list.Item{}

	for _, v := range config.DaysList {
		items = append(items, item(v))
	}

	return items
}

func daysHelpOptions() []key.Binding {
	return []key.Binding{keys.Exec, keys.Read}
}
