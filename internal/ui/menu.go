package ui

import (
	"fmt"
	"strings"

	"gcm/internal/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	items    []model.GitChange
	cursor   int
	selected map[int]bool
	quitting bool
	canceled bool
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
		case " ", "space":
			// toggle selection
			if m.selected[m.cursor] {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}
		case "enter":
			// Confirm selection
			m.quitting = true
			return m, tea.Quit
		case "a":
			// Select all
			for i := range m.items {
				m.selected[i] = true
			}
		case "d":
			// Deselect all
			m.selected = make(map[int]bool)
		case "i":
			// Invert selection
			newSelected := make(map[int]bool)
			for i := range m.items {
				if !m.selected[i] {
					newSelected[i] = true
				}
			}
			m.selected = newSelected
		case "q", "esc", "ctrl+c":
			m.quitting = true
			m.canceled = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	if len(m.items) == 0 {
		return "No changes detected\n"
	}

	var b strings.Builder

	b.WriteString(titleStyle.Render("📂 File Selection") + "\n\n")
	b.WriteString(promptStyle.Render("Navigate with ↑/↓, SPACE to mark, ENTER to continue") + "\n\n")

	// Group by type for display
	categorized := make(map[string][]int)
	for i, item := range m.items {
		typ := item.DisplayType()
		categorized[typ] = append(categorized[typ], i)
	}

	// Display order
	order := []string{"MODIFIED", "ADDED", "DELETED", "RENAMED", "UNTRACKED"}

	for _, typ := range order {
		indices, ok := categorized[typ]
		if !ok || len(indices) == 0 {
			continue
		}

		typeStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))
		b.WriteString(typeStyle.Render(fmt.Sprintf("%s:", typ)) + "\n")

		for _, i := range indices {
			it := m.items[i]
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			checked := " "
			if m.selected[i] {
				checked = "x"
			}

			line := fmt.Sprintf("%s [%s] %s\n", cursor, checked, it.Path)

			if m.cursor == i {
				b.WriteString(infoStyle.Render(line))
			} else {
				b.WriteString(line)
			}
		}
		b.WriteString("\n")
	}

	selectedCount := len(m.selected)
	totalCount := len(m.items)
	b.WriteString(promptStyle.Render(fmt.Sprintf("Selected: %d/%d\n", selectedCount, totalCount)))
	b.WriteString(promptStyle.Render("Shortcuts: 'a' (select all), 'd' (deselect all), 'i' (invert), 'q' (cancel)\n"))
	b.WriteString(promptStyle.Render("Tip: Group related changes in the same commit\n"))

	return b.String()
}

func Run(items []model.GitChange) ([]model.GitChange, error) {
	p := tea.NewProgram(New(items))
	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	modelPtr := m.(*Model)
	if modelPtr.canceled {
		return nil, nil
	}

	var res []model.GitChange
	for i := range modelPtr.items {
		if modelPtr.selected[i] {
			res = append(res, modelPtr.items[i])
		}
	}
	return res, nil
}
