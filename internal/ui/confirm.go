package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ConfirmModel struct {
	prompt   string
	options  string // e.g., "(y/n)"
	result   string
	quitting bool
}

func NewConfirmModel(prompt, options string) *ConfirmModel {
	return &ConfirmModel{
		prompt:  prompt,
		options: options,
	}
}

func (m *ConfirmModel) Init() tea.Cmd {
	return nil
}

func (m *ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y":
			m.result = "y"
			m.quitting = true
			return m, tea.Quit
		case "n":
			m.result = "n"
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			m.result = "n"
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *ConfirmModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s %s ", m.prompt, m.options))
	return b.String()
}

func Confirm(prompt, options string) (bool, error) {
	p := tea.NewProgram(NewConfirmModel(prompt, options))
	m, err := p.Run()
	if err != nil {
		return false, err
	}

	model := m.(*ConfirmModel)
	return model.result == "y", nil
}
