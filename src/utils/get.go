package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

const jiraBaseURL = "https://buk.atlassian.net/rest/api/3/issue/"

func SendGetHttpRequest(ticketID string) ([]byte, error) {
	email := os.Getenv("JIRA_EMAIL")
	token := os.Getenv("JIRA_TOKEN")
	if email == "" || token == "" {
		return nil, fmt.Errorf("JIRA_EMAIL or JIRA_TOKEN environment variable is not set")
	}

	auth := base64.StdEncoding.EncodeToString([]byte(email + ":" + token))

	url := jiraBaseURL + ticketID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get ticket: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
