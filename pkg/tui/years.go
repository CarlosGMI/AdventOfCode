package tui

import (
	"AdventOfCode/pkg/config"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type YearsModel struct {
	list   list.Model
	choice string
}

func NewYearsModel() YearsModel {
	model := YearsModel{}
	items := model.PopulateItems()
	list := createList(items, "Select year:", config.Years, yearsHelpOptions)

	model.list = list

	return model
}

func (model YearsModel) Init() tea.Cmd {
	return nil
}

func (model YearsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				daysModel := newDaysModel(model.choice)

				return daysModel.Update(tea.KeyMsg{})
			}

			return model, tea.Quit
		}
	}

	var cmd tea.Cmd
	model.list, cmd = model.list.Update(msg)

	return model, cmd
}

func (model YearsModel) View() string {
	return fmt.Sprintf("\n%s\n", model.list.View())
}

func (model YearsModel) PopulateItems() []list.Item {
	items := []list.Item{}

	for _, v := range config.YearsList {
		items = append(items, item(v))
	}

	return items
}

func yearsHelpOptions() []key.Binding {
	return []key.Binding{keys.Enter}
}