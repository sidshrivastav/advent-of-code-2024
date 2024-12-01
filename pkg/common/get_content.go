package common

import (
	"fmt"
	"io"
	"os"
)

func GetContent(input_file string) (string, error) {
	file, err := os.Open(input_file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}
