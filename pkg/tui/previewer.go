package tui

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const viewportWidth = 110

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type previewer struct {
	viewport     viewport.Model
	selectedYear string
	selectedDay  string
}

func newPreviewer(selectedYear string, selectedDay string) (previewer, error) {
	model := previewer{
		selectedYear: selectedYear,
		selectedDay:  selectedDay,
	}
	viewport := viewport.New(viewportWidth, 50)
	viewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)
	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(viewportWidth),
	)

	if err != nil {
		return previewer{}, err
	}

	content := model.contentPreview()
	str, err := renderer.Render(content)

	if err != nil {
		return previewer{}, err
	}

	viewport.SetContent(str)
	model.viewport = viewport

	return model, nil
}

func (model previewer) Init() tea.Cmd {
	return nil
}

func (model previewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return model, tea.Quit
		case "esc":
			model := newDaysModel(model.selectedYear, model.selectedDay)

			return model.Update(tea.KeyMsg{})
		default:
			var cmd tea.Cmd
			model.viewport, cmd = model.viewport.Update(msg)
			return model, cmd
		}
	default:
		return model, nil
	}
}

func (model previewer) View() string {
	return model.viewport.View() + model.helpView()
}

func (model previewer) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • esc: go back • q: quit\n")
}

func (model previewer) contentPreview() string {
	content, err := os.ReadFile(fmt.Sprintf("./pkg/app/%s/%s/README.md", model.selectedYear, model.selectedDay))

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
