package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edgarlatorre/bookmark/internal/repositories"
	"os/exec"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, url string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.url }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func initialData() model {
	urls, err := repositories.Read("urls.json")

	if err != nil {
		fmt.Println("Error to read json file:", err)

		m := model{list: list.New(nil, list.NewDefaultDelegate(), 0, 0)}
		m.list.Title = "Bookmark"

		return m
	}

	items := make([]list.Item, len(urls))

	for i, u := range urls {
		items[i] = item{title: u.Name, url: u.Url}
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Bookmark"

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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			if item, ok := m.list.SelectedItem().(item); ok {
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
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
