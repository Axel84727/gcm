package model

import "fmt"

type GitChange struct {
	Index   byte
	Working byte
	Path    string
}

func (g GitChange) StatusKey() string {
	return string([]byte{g.Index, g.Working})
}

func (g GitChange) DisplayLabel() string {
	return fmt.Sprintf("[%c%c] %s", g.Index, g.Working, g.Path)
}

func (g GitChange) DisplayType() string {
	// Categorize based on git status codes
	status := g.StatusKey()
	switch {
	case status == "M " || status == " M" || status == "MM":
		return "MODIFIED"
	case status == " D" || status == "D " || status == "DD":
		return "DELETED"
	case status == "A " || status == " A":
		return "ADDED"
	case status[0] == 'R':
		return "RENAMED"
	case status == "??":
		return "UNTRACKED"
	default:
		return "MODIFIED"
	}
}

type CommitType struct {
	Key         string
	Description string
}

var CommitTypes = []CommitType{
	{"feat", "New feature"},
	{"fix", "Bug fix"},
	{"docs", "Documentation"},
	{"style", "Formatting/styles (doesn't affect logic)"},
	{"refactor", "Refactoring (no functional changes)"},
	{"perf", "Performance improvement"},
	{"test", "Tests"},
	{"chore", "Maintenance tasks"},
	{"build", "Build system"},
	{"ci", "CI/CD"},
}

type CommitInfo struct {
	Type        string
	Title       string
	Description string
}

func (c CommitInfo) FullMessage() string {
	if c.Description != "" {
		return fmt.Sprintf("%s: %s\n\n%s", c.Type, c.Title, c.Description)
	}
	return fmt.Sprintf("%s: %s", c.Type, c.Title)
}
