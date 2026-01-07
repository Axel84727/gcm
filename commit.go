package main

import (
	"bufio"
	"strings"
)

func parseChangedFiles(output string) []GitChange {
	var changes []GitChange

	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			continue
		}

		change := GitChange{
			Index:   line[0],
			Working: line[1],
			Path:    strings.TrimSpace(line[2:]),
		}

		changes = append(changes, change)
	}

	return changes
}
