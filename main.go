package main

import "fmt"

func main() {
	output, err := checkChangedFiles()
	if err != nil {
		fmt.Println("error checking changed files:", err)
		return
	}

	changes := parseChangedFiles(output)

	if len(changes) == 0 {
		fmt.Println("no changes detected")
		return
	}

	for _, c := range changes {
		fmt.Printf("[%c%c] %s\n", c.Index, c.Working, c.Path)
	}
}
