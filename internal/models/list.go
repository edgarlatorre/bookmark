package models

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edgarlatorre/bookmark/internal/repositories"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type keymap struct {
	Create key.Binding
}

var Keymap = keymap{
	Create: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "New url"),
	),
}

type item struct {
	title, url string
}

type ListModel struct {
	model list.Model
}

func NewListModel() list.Model {
	repository := repositories.UrlRepository{FilePath: "urls.csv"}
	urls, err := repository.Read()
	m := list.Model{}

	if err != nil {
		fmt.Println("Error to read csv file:", err)

		return m
	}

	items := make([]list.Item, len(urls))

	for i, u := range urls {
		items[i] = u
	}

	m = list.New(items, list.NewDefaultDelegate(), 0, 0)
	m.Title = "Bookmark"
	m.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			Keymap.Create,
		}
	}

	return m
}

func UpdateList(m list.Model, msg tea.Msg) (list.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "c":
			return m, tea.Quit
		case "enter", " ":
			if item, ok := m.SelectedItem().(repositories.Url); ok {
				cmd := exec.Command("open", item.Description())
				_, err := cmd.Output()

				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Println("Not found")
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m, cmd = m.Update(msg)

	return m, cmd
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) View() string {
	return m.model.View()
}
