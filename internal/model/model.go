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
