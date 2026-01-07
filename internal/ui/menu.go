package ui

import (
	"fmt"
	"strings"

	"gcm/internal/model"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	items    []model.GitChange
	cursor   int
	selected map[int]bool
	quitting bool
}

func New(items []model.GitChange) *Model {
	return &Model{
		items:    items,
		cursor:   0,
		selected: make(map[int]bool),
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			// toggle selection
			if m.selected[m.cursor] {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}
		case "c":
			m.quitting = true
			return m, tea.Quit
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	if len(m.items) == 0 {
		return "no changes detected\n"
	}

	var b strings.Builder
	b.WriteString("Use ↑/↓ to move, Enter to toggle, c to confirm, q to quit\n\n")

	for i, it := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected[i] {
			checked = "x"
		}

		line := fmt.Sprintf("%s [%s] %s\n", cursor, checked, it.DisplayLabel())
		b.WriteString(line)
	}

	if m.quitting {
		b.WriteString("\nQuitting...\n")
	}

	return b.String()
}

func Run(items []model.GitChange) ([]model.GitChange, error) {
	p := tea.NewProgram(New(items))
	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	modelPtr := m.(*Model)
	var res []model.GitChange
	for i := range modelPtr.items {
		if modelPtr.selected[i] {
			res = append(res, modelPtr.items[i])
		}
	}
	return res, nil
}
