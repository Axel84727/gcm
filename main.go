package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gcm/internal/changes"
	gitpkg "gcm/internal/git"
	"gcm/internal/ui"
)

func main() {
	output, err := gitpkg.CheckChangedFiles()
	if err != nil {
		fmt.Println("error checking changed files:", err)
		return
	}

	changesList := changes.ParseChangedFiles(output)

	if len(changesList) == 0 {
		fmt.Println("no changes detected")
		return
	}

	selected, err := ui.Run(changesList)
	if err != nil {
		fmt.Println("error running UI:", err)
		return
	}

	if len(selected) == 0 {
		fmt.Println("no files selected, exiting")
		return
	}

	cats := changes.CategorizeByStatus(selected)
	fmt.Println("Selected grouped by status:")
	for k, v := range cats {
		fmt.Printf("%s:\n", k)
		for _, it := range v {
			fmt.Println("  -", it.Path)
		}
	}

	var paths []string
	for _, it := range selected {
		paths = append(paths, it.Path)
	}

	err = gitpkg.Add(paths)
	if err != nil {
		fmt.Println("error during git add:", err)
		return
	}

	fmt.Print("Commit message: ")
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	if msg == "" {
		msg = "WIP"
	}

	err = gitpkg.Commit(msg)
	if err != nil {
		fmt.Println("error during git commit:", err)
		return
	}

	fmt.Println("Done.")
}
