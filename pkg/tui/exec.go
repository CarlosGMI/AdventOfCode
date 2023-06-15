package tui

import (
	"AdventOfCode/pkg/app"
	"AdventOfCode/pkg/config"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func waitForActivity(sub chan app.ChallengeMsg) tea.Cmd {
	return func() tea.Msg {
		return app.ChallengeMsg(<-sub)
	}
}

type execModel struct {
	sub          chan app.ChallengeMsg
	data         []string
	selectedYear string
	selectedDay  string
}

func newExecModel(year string, day string) execModel {
	return execModel{
		sub:          make(chan app.ChallengeMsg),
		selectedYear: year,
		selectedDay:  day,
	}
}

func (model execModel) Init() tea.Cmd {
	return tea.Batch(
		model.listenForActivity(),
		waitForActivity(model.sub),
	)
}

func (model execModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return model, tea.Quit
		case "enter":
			model.emptySub()
			model.data = nil

			return model, model.listenForActivity()
		case "esc":
			model.emptySub()
			model := newDaysModel(model.selectedYear, model.selectedDay)

			return model.Update(tea.KeyMsg{})
		default:
			return model, nil
		}
	case app.ChallengeMsg:
		model.data = append(model.data, msg.Data)

		return model, waitForActivity(model.sub)
	default:
		return model, nil
	}
}

func (model execModel) View() string {
	var style = lipgloss.NewStyle().
		PaddingLeft(4)

	return style.Render(fmt.Sprintf("\n%s\n\n\n\n\n\n\n\n\n\n\n\n%s", strings.Join(model.data, "\n"), model.help()))
}

func (model execModel) help() string {
	return helpStyle("esc: go back • enter: re-execute • q: quit\n\n")
}

func (model execModel) listenForActivity() tea.Cmd {
	return func() tea.Msg {
		challenge := config.Challenges[fmt.Sprintf("%s-%s", model.selectedYear, model.selectedDay)]

		if challenge == nil {
			errorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(config.ErrorColor)).Render
			model.sub <- app.ChallengeMsg{
				Data: errorStyle("Execution function not found for this challenge"),
			}

			return nil
		} else {
			challenge.Exec(&model.sub)
		}

		return app.ChallengeMsg(<-model.sub)
	}
}

func (model execModel) emptySub() {
	for len(model.sub) > 0 {
		<-model.sub
	}
}
