package utils

import (
	"fmt"
	"os"
	"time"
)

const outputFile = "diff_output.txt"

func AppendToFile(content string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	data := fmt.Sprintf("\n--- %s ---\n%s\n", timestamp, content)

	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	return err
}