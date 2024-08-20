package utils

import (
	"bytes"
	"io"
	"net/http"
)

const apiEndpoint = "https://api.ejemplo.com/endpoint"

func SendHttpPost(content string) error {
	payload := bytes.NewBufferString(content)
	resp, err := http.Post(apiEndpoint, "text/plain", payload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	return err
}