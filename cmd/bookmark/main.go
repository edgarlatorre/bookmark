package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edgarlatorre/bookmark/internal/models"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type State uint

const (
	listView State = iota
	formView
)

type keymap struct {
	Create key.Binding
}

var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "New url"),
	),
}

type model struct {
	state State
	list  list.Model
	form  models.FormModel
}

func initialData() model {
	m := model{
		state: listView,
		list:  models.NewListModel(),
		form:  models.NewFormModel(),
	}

	m.list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			Keymap.Create,
		}
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialData(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.state == listView {
		m.list, cmd = models.UpdateList(m.list, msg)
	} else {
		m.form, cmd = models.UpdateForm(m.form, msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			m.state = formView
		case "esc", "back":
			m.state = listView
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.state == listView {
		return m.list.View()
	} else {
		return m.form.View()
	}
}
