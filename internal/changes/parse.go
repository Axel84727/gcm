package changes

import (
	"bufio"
	"strings"

	"gcm/internal/model"
)

func ParseChangedFiles(output string) []model.GitChange {
	var changes []model.GitChange

	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			continue
		}

		change := model.GitChange{
			Index:   line[0],
			Working: line[1],
			Path:    strings.TrimSpace(line[2:]),
		}

		changes = append(changes, change)
	}

	return changes
}
