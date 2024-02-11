package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/edgarlatorre/bookmark/internal/repositories"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type FormModel struct {
	focusIndex int
	form       []textinput.Model
	cursorMode cursor.Mode
}

func NewFormModel() FormModel {
	nameInput := textinput.New()
	nameInput.Placeholder = "Name"
	nameInput.Focus()

	urlInput := textinput.New()
	urlInput.Placeholder = "Url"

	m := FormModel{
		form: []textinput.Model{nameInput, urlInput},
	}

	return m
}

func (m FormModel) Init() tea.Cmd {
	return textinput.Blink
}

func UpdateForm(m FormModel, msg tea.Msg) (FormModel, tea.Cmd) {
	cmd := m.updateInputs(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "shift+tab", "enter":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.form) {
				if m.form[0].Value() != "" && m.form[0].Value() != "" {
					r := repositories.UrlRepository{FilePath: "urls.csv"}
					_, err := r.Create(m.form[1].Value(), m.form[0].Value())

					if err != nil {
						return m, tea.Quit
					}
				}
				return m, tea.ClearScreen
			}

			if s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.form) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.form)
			}

			cmds := make([]tea.Cmd, len(m.form))
			for i := 0; i <= len(m.form)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.form[i].Focus()
					m.form[i].PromptStyle = focusedStyle
					m.form[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.form[i].Blur()
				m.form[i].PromptStyle = noStyle
				m.form[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	return m, cmd
}

func (m *FormModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.form))

	for i := range m.form {
		m.form[i], cmds[i] = m.form[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m FormModel) View() string {
	var b strings.Builder
	inputStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)
	b.WriteString("  " + inputStyle.Render("Bookmark"))
	b.WriteString("\n\n")

	for i := range m.form {
		b.WriteString("  " + m.form[i].View() + "\n\n")
	}

	button := &blurredButton

	if m.focusIndex == len(m.form) {
		button = &focusedButton
	}

	fmt.Fprintf(&b, "  %s\n\n", *button)

	return b.String()
}
