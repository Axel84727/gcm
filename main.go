package main

import "fmt"

func main() {
	changes, err := checkChangedFiles()
	if err != nil {
		fmt.Println("Error checking changed files")
		return
	}
	if changes == "" {
		fmt.Println("not changes detected")
		return
	}
	fmt.Println(changes)
}
