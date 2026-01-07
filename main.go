package main

import (
	"fmt"
	"os"
	"strings"

	"gcm/internal/changes"
	gitpkg "gcm/internal/git"
	"gcm/internal/ui"
	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	warningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("226"))
)

func main() {
	// Step 1: Check for changes
	output, err := gitpkg.CheckChangedFiles()
	if err != nil {
		fmt.Println("❌ Error checking changed files:", err)
		os.Exit(1)
	}

	changesList := changes.ParseChangedFiles(output)

	if len(changesList) == 0 {
		fmt.Println(successStyle.Render("✨ Working tree clean, nothing to commit"))
		return
	}

	// Step 2: Branch management
	currentBranch, err := gitpkg.GetCurrentBranch()
	if err != nil {
		fmt.Println("❌ Error getting current branch:", err)
		os.Exit(1)
	}

	isMainBranch := gitpkg.IsMainBranch(currentBranch)

	branchName, confirmed, err := ui.RunBranchSelection(currentBranch, isMainBranch)
	if err != nil {
		fmt.Println("❌ Error in branch selection:", err)
		os.Exit(1)
	}

	if !confirmed {
		fmt.Println("Canceled.")
		return
	}

	// Create new branch if needed
	if branchName != currentBranch {
		err = gitpkg.CreateBranch(branchName)
		if err != nil {
			fmt.Println("❌ Error creating branch:", err)
			os.Exit(1)
		}
		fmt.Println(successStyle.Render(fmt.Sprintf("✓ Created and switched to branch '%s'", branchName)))
	} else {
		fmt.Println(infoStyle.Render(fmt.Sprintf("📌 Using branch '%s'", branchName)))
	}

	// Track commits created in this session
	var commitsCreated []string

	// Step 3: Loop until no more changes or user quits
	for {
		// Re-check for changes
		output, err := gitpkg.CheckChangedFiles()
		if err != nil {
			fmt.Println("❌ Error checking changed files:", err)
			break
		}

		changesList = changes.ParseChangedFiles(output)

		if len(changesList) == 0 {
			fmt.Println(successStyle.Render("\n✨ All files committed!"))
			break
		}

		fmt.Printf("\n%s\n", infoStyle.Render(fmt.Sprintf("📋 %d file(s) with changes", len(changesList))))

		// Step 4: File selection
		selected, err := ui.Run(changesList)
		if err != nil {
			fmt.Println("❌ Error running UI:", err)
			break
		}

		if selected == nil || len(selected) == 0 {
			fmt.Println("No files selected, exiting.")
			break
		}

		// Show warning for large commits
		if len(selected) > 10 {
			fmt.Println(warningStyle.Render(fmt.Sprintf("⚠️  Warning: Large commit with %d files", len(selected))))
		}

		// Step 5: Commit type selection
		commitType, isCustom, err := ui.RunCommitTypeSelection()
		if err != nil {
			fmt.Println("❌ Error selecting commit type:", err)
			break
		}

		if commitType == "" && !isCustom {
			fmt.Println("No commit type selected, exiting.")
			break
		}

		// Handle custom commit type
		if isCustom {
			customType, ok, err := ui.GetInput("Enter custom commit type:")
			if err != nil {
				fmt.Println("❌ Error getting custom type:", err)
				break
			}
			if !ok || customType == "" {
				fmt.Println("No custom type provided, exiting.")
				break
			}
			commitType = customType
		}

		// Step 6: Commit message
		title, description, confirmed, err := ui.RunCommitMessage(commitType)
		if err != nil {
			fmt.Println("❌ Error getting commit message:", err)
			break
		}

		if !confirmed {
			fmt.Println("Commit message not confirmed, exiting.")
			break
		}

		// Step 7: Stage files
		var paths []string
		for _, it := range selected {
			paths = append(paths, it.Path)
		}

		err = gitpkg.Add(paths)
		if err != nil {
			fmt.Println("❌ Error during git add:", err)
			break
		}

		// Step 8: Commit
		fullMessage := fmt.Sprintf("%s: %s", commitType, title)
		if description != "" {
			fullMessage = fmt.Sprintf("%s\n\n%s", fullMessage, description)
		}

		err = gitpkg.Commit(fullMessage)
		if err != nil {
			fmt.Println("❌ Error during git commit:", err)
			break
		}

		fmt.Println(successStyle.Render(fmt.Sprintf("✓ Commit created: [%s] %s", commitType, title)))
		commitsCreated = append(commitsCreated, fmt.Sprintf("[%s] %s", commitType, title))

		// Step 9: Check if there are more uncommitted files
		output, err = gitpkg.CheckChangedFiles()
		if err != nil {
			break
		}

		remainingChanges := changes.ParseChangedFiles(output)

		if len(remainingChanges) == 0 {
			fmt.Println(successStyle.Render("\n✨ All files committed!"))
			break
		}

		// Ask if user wants to continue
		fmt.Printf("\n%s\n", infoStyle.Render(fmt.Sprintf("Current status:")))
		fmt.Printf("  - %d file(s) committed\n", len(selected))
		fmt.Printf("  - %d file(s) still uncommitted\n\n", len(remainingChanges))

		continueCommit, err := ui.Confirm("Want to create another commit on the same branch?", "(y/n)")
		if err != nil || !continueCommit {
			break
		}
	}

	// Step 10: Show summary and offer push
	if len(commitsCreated) > 0 {
		fmt.Println("\n" + strings.Repeat("─", 50))
		fmt.Println(successStyle.Render("📦 Commits created in this session:"))
		for i, commit := range commitsCreated {
			fmt.Printf("  %d. %s\n", i+1, commit)
		}
		fmt.Println(strings.Repeat("─", 50))

		// Offer to push
		shouldPush, err := ui.Confirm("\nPush to remote?", "(y/n)")
		if err == nil && shouldPush {
			// Check if remote branch exists
			hasRemote, _ := gitpkg.HasRemoteBranch(branchName)

			fmt.Printf("\n🚀 Pushing to origin/%s...\n", branchName)

			err = gitpkg.Push(branchName, !hasRemote)
			if err != nil {
				fmt.Println("❌ Error during push:", err)
			} else {
				fmt.Println(successStyle.Render("✓ Successfully pushed to remote"))
			}
		} else {
			fmt.Println(infoStyle.Render("\n💡 You can push later with: git push origin " + branchName))
		}
	}

	fmt.Println("\n" + successStyle.Render("✨ Done."))
}
