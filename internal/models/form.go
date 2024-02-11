package models

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	t := textinput.New()
	t.Placeholder = "Url"
	t.Focus()

	m := FormModel{
		form: []textinput.Model{t},
	}

	return m
}

func (m FormModel) Init() tea.Cmd {
	return textinput.Blink
}

func UpdateForm(m FormModel, msg tea.Msg) (FormModel, tea.Cmd) {
	cmd := m.updateInputs(msg)

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
		b.WriteString("  " + m.form[i].View())
	}

	button := &blurredButton
	fmt.Fprintf(&b, "\n\n  %s\n\n", *button)

	return b.String()
}
