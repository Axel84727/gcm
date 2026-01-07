package ui

import (
	"fmt"
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type CommitMessageModel struct {
	commitType  string
	title       string
	description string
	mode        string // "title" or "description" or "preview"
	err         string
	quitting    bool
	confirmed   bool
}

func NewCommitMessageModel(commitType string) *CommitMessageModel {
	return &CommitMessageModel{
		commitType: commitType,
		mode:       "title",
	}
}

func (m *CommitMessageModel) Init() tea.Cmd {
	return nil
}

func (m *CommitMessageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "esc":
			if m.mode == "description" || m.mode == "preview" {
				m.mode = "title"
				m.err = ""
			} else {
				m.quitting = true
				return m, tea.Quit
			}

		case "enter":
			if m.mode == "title" {
				if err := validateTitle(m.title); err != nil {
					m.err = err.Error()
					return m, nil
				}
				m.mode = "description"
				m.err = ""
			} else if m.mode == "description" {
				// Empty enter skips description
				m.mode = "preview"
				m.err = ""
			} else if m.mode == "preview" {
				// Confirm
				m.confirmed = true
				m.quitting = true
				return m, tea.Quit
			}

		case "ctrl+d":
			if m.mode == "description" {
				m.mode = "preview"
				m.err = ""
			}

		case "backspace":
			if m.mode == "title" && len(m.title) > 0 {
				m.title = m.title[:len(m.title)-1]
				m.err = ""
			} else if m.mode == "description" && len(m.description) > 0 {
				m.description = m.description[:len(m.description)-1]
			}

		case "y":
			if m.mode == "preview" {
				m.confirmed = true
				m.quitting = true
				return m, tea.Quit
			}

		case "n":
			if m.mode == "preview" {
				m.mode = "title"
				m.err = ""
			}

		case "e":
			if m.mode == "preview" {
				m.mode = "title"
				m.err = ""
			}

		default:
			if m.mode == "title" && len(msg.String()) == 1 {
				m.title += msg.String()
				m.err = ""
			} else if m.mode == "description" {
				if len(msg.String()) == 1 {
					m.description += msg.String()
				}
			}
		}
	}

	return m, nil
}

func (m *CommitMessageModel) View() string {
	if m.quitting && m.confirmed {
		return ""
	}

	var b strings.Builder

	if m.mode == "title" {
		b.WriteString(titleStyle.Render("📝 Commit Title") + "\n\n")
		b.WriteString(fmt.Sprintf("Type: %s\n\n", infoStyle.Render(m.commitType)))
		b.WriteString("Commit title (required):\n")
		b.WriteString("> " + m.title + "_\n\n")

		if m.err != "" {
			b.WriteString(errorStyle.Render("❌ "+m.err) + "\n\n")
		}

		titleLen := len(m.title)
		if titleLen > 50 && titleLen <= 72 {
			b.WriteString(promptStyle.Render(fmt.Sprintf("⚠️  Warning: Title is %d characters (recommended max: 50)\n", titleLen)))
		}

		b.WriteString(promptStyle.Render("Rules:\n"))
		b.WriteString(promptStyle.Render("  - Min 10 characters, max 72\n"))
		b.WriteString(promptStyle.Render("  - First letter lowercase (conventional commits)\n"))
		b.WriteString(promptStyle.Render("  - No period at the end\n"))
		b.WriteString(promptStyle.Render("  - Concise and clear description\n\n"))
		b.WriteString(promptStyle.Render("Example: \"add email validation in registration form\"\n"))
		b.WriteString(promptStyle.Render("Press Enter to continue\n"))

	} else if m.mode == "description" {
		b.WriteString(titleStyle.Render("📝 Detailed Description (Optional)") + "\n\n")
		b.WriteString(fmt.Sprintf("%s: %s\n\n", m.commitType, m.title))
		b.WriteString("Description:\n")

		if m.description == "" {
			b.WriteString(promptStyle.Render("(Press Enter to skip, or type description)\n\n"))
		} else {
			b.WriteString(m.description + "_\n\n")
		}

		b.WriteString(promptStyle.Render("Tip: Explain the 'why', not the 'what' (that's in the diff)\n"))
		b.WriteString(promptStyle.Render("Press Ctrl+D or Enter (empty) to finish, Esc to go back\n"))

	} else if m.mode == "preview" {
		b.WriteString(titleStyle.Render("📋 Commit Preview") + "\n\n")
		b.WriteString("+" + strings.Repeat("-", 70) + "+\n")
		b.WriteString("|" + strings.Repeat(" ", 70) + "|\n")
		b.WriteString(fmt.Sprintf("|  %s: %-63s|\n", m.commitType, m.title))
		b.WriteString("|" + strings.Repeat(" ", 70) + "|\n")

		if m.description != "" {
			// Split description into lines
			descLines := splitIntoLines(m.description, 66)
			for _, line := range descLines {
				b.WriteString(fmt.Sprintf("|  %-68s|\n", line))
			}
			b.WriteString("|" + strings.Repeat(" ", 70) + "|\n")
		}

		b.WriteString("+" + strings.Repeat("-", 70) + "+\n\n")
		b.WriteString("Confirm commit message? (y/n/e to edit): ")
	}

	return b.String()
}

func validateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if len(title) < 10 {
		return fmt.Errorf("title too short (min 10 characters)")
	}

	if len(title) > 72 {
		return fmt.Errorf("title too long (max 72 characters)")
	}

	if strings.HasSuffix(title, ".") {
		return fmt.Errorf("title should not end with a period")
	}

	if len(title) > 0 && unicode.IsUpper(rune(title[0])) {
		return fmt.Errorf("title should start with lowercase letter (conventional commits)")
	}

	return nil
}

func splitIntoLines(text string, maxWidth int) []string {
	var lines []string
	words := strings.Fields(text)

	if len(words) == 0 {
		return lines
	}

	currentLine := words[0]
	for _, word := range words[1:] {
		if len(currentLine)+1+len(word) <= maxWidth {
			currentLine += " " + word
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return lines
}

func RunCommitMessage(commitType string) (string, string, bool, error) {
	p := tea.NewProgram(NewCommitMessageModel(commitType))
	m, err := p.Run()
	if err != nil {
		return "", "", false, err
	}

	model := m.(*CommitMessageModel)
	if !model.confirmed {
		return "", "", false, nil
	}

	return model.title, model.description, true, nil
}
