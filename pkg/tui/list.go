package tui

import (
	"AdventOfCode/pkg/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const listWidth = 20
const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(config.SelectedColor))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	listHelpStyle     = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type shortHelpItems func() []key.Binding

func createList(items []list.Item, title string, listType string, help shortHelpItems) list.Model {
	list := list.New(items, itemDelegate{}, listWidth, listHeight)
	list.Title = title
	list.Styles.Title = titleStyle
	list.Styles.PaginationStyle = paginationStyle
	list.Styles.HelpStyle = listHelpStyle
	list.AdditionalShortHelpKeys = help

	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(false)

	return list
}
