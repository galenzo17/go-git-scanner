package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetJiraTicketID() (string, error) {
	log.Println("starting to get the current branch name...")
	gitDir := os.Getenv("GIT_DIR")
	if gitDir == "" {
		return "", fmt.Errorf("GIT_DIR environment variable is not set")
	}
	cmd := exec.Command("git", "-C", gitDir, "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %v", err)
	}

	branchName := strings.TrimSpace(string(output))
	log.Printf("current branch name: %s\n", branchName)

	// Split the branch name by "/"
	parts := strings.Split(branchName, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("branch name format is incorrect, unable to extract Jira ticket ID")
	}

	// The Jira ticket ID is assumed to be the second part
	jiraTicketID := parts[1]
	log.Printf("extracted Jira ticket ID: %s\n", jiraTicketID)

	return jiraTicketID, nil
}
