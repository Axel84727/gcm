package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	prompt   string
	value    string
	quitting bool
	canceled bool
}

func NewInputModel(prompt string) *InputModel {
	return &InputModel{
		prompt: prompt,
	}
}

func (m *InputModel) Init() tea.Cmd {
	return nil
}

func (m *InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c", "esc":
			m.canceled = true
			m.quitting = true
			return m, tea.Quit
		case "backspace":
			if len(m.value) > 0 {
				m.value = m.value[:len(m.value)-1]
			}
		default:
			if len(msg.String()) == 1 {
				m.value += msg.String()
			}
		}
	}
	return m, nil
}

func (m *InputModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString(titleStyle.Render("✏️  Custom Commit Type") + "\n\n")
	b.WriteString(m.prompt + "\n")
	b.WriteString("> " + m.value + "_\n\n")
	b.WriteString(promptStyle.Render("Press Enter to confirm, Esc to cancel\n"))
	return b.String()
}

func GetInput(prompt string) (string, bool, error) {
	p := tea.NewProgram(NewInputModel(prompt))
	m, err := p.Run()
	if err != nil {
		return "", false, err
	}

	model := m.(*InputModel)
	if model.canceled {
		return "", false, nil
	}

	return strings.TrimSpace(model.value), true, nil
}
