package tui

import (
	"AdventOfCode/pkg/config"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type DaysModel struct {
	list         list.Model
	choice       string
	selectedYear string
}

func newDaysModel(year string, day string) DaysModel {
	model := DaysModel{selectedYear: year}
	items := model.PopulateItems()
	list := createList(items, "Select day:", config.Days, daysHelpOptions)
	model.list = list

	if day != "" {
		currentDayIndex, err := strconv.Atoi(day)

		if err != nil {
			log.Fatal(err)
		}

		model.list.Select(currentDayIndex - 1)
	}

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

			newModel := newExecModel(model.selectedYear, model.choice)

			return newModel, newModel.Init()
		}

		switch {
		case key.Matches(msg, keys.Read):
			i, ok := model.list.SelectedItem().(item)

			if ok {
				model.choice = string(i)
			}

			preview, err := newPreviewer(model.selectedYear, model.choice)

			if err != nil {
				log.Fatal(err)
			}

			return preview.Update(tea.KeyMsg{})

		case key.Matches(msg, keys.Back):
			newModel := NewYearsModel(model.selectedYear)

			return newModel.Update(tea.KeyMsg{})
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

	entries, err := os.ReadDir(fmt.Sprintf("./pkg/app/instructions/%s", model.selectedYear))

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range entries {
		items = append(items, item(v.Name()))
	}

	return items
}

func daysHelpOptions() []key.Binding {
	return []key.Binding{keys.Back, keys.Exec, keys.Read}
}
