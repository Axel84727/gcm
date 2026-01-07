package ui

import (
	"fmt"
	"strings"

	"gcm/internal/model"

	tea "github.com/charmbracelet/bubbletea"
)

type CommitTypeModel struct {
	types    []model.CommitType
	cursor   int
	selected string
	custom   bool
	input    string
	quitting bool
}

func NewCommitTypeModel() *CommitTypeModel {
	return &CommitTypeModel{
		types:  model.CommitTypes,
		cursor: 0,
	}
}

func (m *CommitTypeModel) Init() tea.Cmd {
	return nil
}

func (m *CommitTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.types)-1 {
				m.cursor++
			}

		case "enter":
			m.selected = m.types[m.cursor].Key
			m.quitting = true
			return m, tea.Quit

		case "c":
			// Custom type
			m.custom = true
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *CommitTypeModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	b.WriteString(titleStyle.Render("📝 Select Commit Type") + "\n\n")

	for i, t := range m.types {
		cursor := "  "
		if m.cursor == i {
			cursor = "> "
		}

		line := fmt.Sprintf("%s%d. %-10s - %s\n", cursor, i+1, t.Key, t.Description)
		if m.cursor == i {
			b.WriteString(infoStyle.Render(line))
		} else {
			b.WriteString(line)
		}
	}

	b.WriteString("\n")
	b.WriteString(promptStyle.Render("Press 'c' for custom type, Enter to select, q to quit\n"))

	return b.String()
}

func RunCommitTypeSelection() (string, bool, error) {
	p := tea.NewProgram(NewCommitTypeModel())
	m, err := p.Run()
	if err != nil {
		return "", false, err
	}

	resultModel := m.(*CommitTypeModel)
	if resultModel.selected == "" && !resultModel.custom {
		return "", false, nil
	}

	return resultModel.selected, resultModel.custom, nil
}
