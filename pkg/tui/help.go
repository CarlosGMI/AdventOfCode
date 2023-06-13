package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Enter key.Binding
	Exec  key.Binding
	Read  key.Binding
}

var keys = keyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Exec: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "execute"),
	),
	Read: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "read"),
	),
}
