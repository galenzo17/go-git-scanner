package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)


func GetGitDiff() (string, error) {
	log.Println("starting git diff....")
	gitDir := os.Getenv("GIT_DIR")
	if gitDir == "" {
		return "", fmt.Errorf("GIT_DIR environment variable is not set")
	}
	cmd := exec.Command("git", "-C", gitDir, "diff", "master")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	log.Println("succes on obtainig git diff....")
	return string(output), nil
}