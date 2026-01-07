package git

import (
	"os"
	"os/exec"
)

func CheckChangedFiles() (string, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
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
