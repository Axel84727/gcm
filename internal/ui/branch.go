package ui

import (
	"fmt"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))
)

type BranchModel struct {
	currentBranch string
	isMainBranch  bool
	input         string
	cursor        int
	mode          string // "confirm" or "input"
	err           string
	quitting      bool
	confirmed     bool
	newBranch     string
}

func NewBranchModel(currentBranch string, isMainBranch bool) *BranchModel {
	mode := "input"
	if !isMainBranch {
		mode = "confirm"
	}

	return &BranchModel{
		currentBranch: currentBranch,
		isMainBranch:  isMainBranch,
		mode:          mode,
		input:         "",
	}
}

func (m *BranchModel) Init() tea.Cmd {
	return nil
}

func (m *BranchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			if m.mode == "confirm" {
				// User confirmed using current branch
				m.confirmed = true
				m.newBranch = m.currentBranch
				m.quitting = true
				return m, tea.Quit
			} else if m.mode == "input" {
				// Validate branch name
				if err := validateBranchName(m.input); err != nil {
					m.err = err.Error()
					return m, nil
				}
				m.confirmed = true
				m.newBranch = m.input
				m.quitting = true
				return m, tea.Quit
			}

		case "n":
			if m.mode == "confirm" {
				m.mode = "input"
				m.input = ""
				m.err = ""
			}

		case "y":
			if m.mode == "confirm" {
				m.confirmed = true
				m.newBranch = m.currentBranch
				m.quitting = true
				return m, tea.Quit
			}

		case "backspace":
			if m.mode == "input" && len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
				m.err = ""
			}

		default:
			if m.mode == "input" && len(msg.String()) == 1 {
				m.input += msg.String()
				m.err = ""
			}
		}
	}

	return m, nil
}

func (m *BranchModel) View() string {
	if m.quitting && m.confirmed {
		return ""
	}

	var b strings.Builder

	if m.isMainBranch {
		b.WriteString(titleStyle.Render("⚠️  Branch Management") + "\n\n")
		b.WriteString(errorStyle.Render(fmt.Sprintf("Cannot commit directly to '%s'.", m.currentBranch)) + "\n")
		b.WriteString("Please create a new branch:\n\n")
	} else if m.mode == "confirm" {
		b.WriteString(titleStyle.Render("📌 Branch Management") + "\n\n")
		b.WriteString(fmt.Sprintf("Current branch: %s\n\n", infoStyle.Render(m.currentBranch)))
		b.WriteString(fmt.Sprintf("Use current branch '%s'? (y/n): ", m.currentBranch))
		return b.String()
	} else {
		b.WriteString(titleStyle.Render("📌 Create New Branch") + "\n\n")
	}

	if m.mode == "input" {
		b.WriteString("Branch name: ")
		b.WriteString(m.input)
		b.WriteString("_\n\n")

		if m.err != "" {
			b.WriteString(errorStyle.Render("❌ "+m.err) + "\n\n")
		}

		b.WriteString(promptStyle.Render("Suggested format: type/short-description\n"))
		b.WriteString(promptStyle.Render("Examples: feat/login, fix/button-crash, chore/deps\n"))
		b.WriteString(promptStyle.Render("Press Enter to confirm, Esc to cancel\n"))
	}

	return b.String()
}

func validateBranchName(name string) error {
	if name == "" {
		return fmt.Errorf("branch name cannot be empty")
	}

	if len(name) > 50 {
		return fmt.Errorf("branch name too long (max 50 characters)")
	}

	// Check for spaces
	if strings.Contains(name, " ") {
		return fmt.Errorf("branch name cannot contain spaces")
	}

	// Check for invalid characters
	invalidChars := regexp.MustCompile(`[~^:?*\[\]\\]`)
	if invalidChars.MatchString(name) {
		return fmt.Errorf("branch name contains invalid characters")
	}

	// Cannot start with dot or slash
	if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "/") {
		return fmt.Errorf("branch name cannot start with '.' or '/'")
	}

	// Cannot end with slash or .lock
	if strings.HasSuffix(name, "/") || strings.HasSuffix(name, ".lock") {
		return fmt.Errorf("invalid branch name ending")
	}

	return nil
}

func RunBranchSelection(currentBranch string, isMainBranch bool) (string, bool, error) {
	p := tea.NewProgram(NewBranchModel(currentBranch, isMainBranch))
	m, err := p.Run()
	if err != nil {
		return "", false, err
	}

	model := m.(*BranchModel)
	if !model.confirmed {
		return "", false, nil
	}

	return model.newBranch, true, nil
}
