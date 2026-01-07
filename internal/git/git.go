package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckChangedFiles() (string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func IsMainBranch(branch string) bool {
	return branch == "main" || branch == "master"
}

func CreateBranch(branchName string) error {
	cmd := exec.Command("git", "checkout", "-b", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CheckoutBranch(branchName string) error {
	cmd := exec.Command("git", "checkout", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Add(paths []string) error {
	if len(paths) == 0 {
		return nil
	}
	args := append([]string{"add"}, paths...)
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Commit(msg string) error {
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CommitWithDescription(title, description string) error {
	var msg string
	if description != "" {
		msg = fmt.Sprintf("%s\n\n%s", title, description)
	} else {
		msg = title
	}
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func HasRemoteBranch(branchName string) (bool, error) {
	cmd := exec.Command("git", "rev-parse", "--verify", fmt.Sprintf("origin/%s", branchName))
	err := cmd.Run()
	if err != nil {
		return false, nil
	}
	return true, nil
}

func Push(branchName string, setUpstream bool) error {
	var cmd *exec.Cmd
	if setUpstream {
		cmd = exec.Command("git", "push", "-u", "origin", branchName)
	} else {
		cmd = exec.Command("git", "push", "origin", branchName)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
