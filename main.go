package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"path/filepath"
    "regexp"

	"github.com/common-nighthawk/go-figure"
)

type Config struct {
	Message string `json:"message"`
}

func validateAndSanitizeFilePath(userInput string) (string, error) {
	// Define a regular expression pattern for allowed file path characters
	pattern := `^[a-zA-Z0-9_\-./]+$`
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(userInput) {
		return "", fmt.Errorf("invalid file path: %s", userInput)
	}
	cleanPath := filepath.Clean(userInput)
	return cleanPath, nil
}

func readConfigFile(filePath string) (*Config, error) {
	// Validate and sanitize the file path
	cleanPath, err := validateAndSanitizeFilePath(filePath)
	if err != nil {
		return nil, err
	}

	// Read the file
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return &config, nil
}


func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file path as an argument")
	}
	filePath := os.Args[1]

	config, err := readConfigFile(filePath)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	myFigure := figure.NewFigure(config.Message, "", true)
	myFigure.Print()

	gi, _ := GetInfo()
	gi.VarDump()
}
