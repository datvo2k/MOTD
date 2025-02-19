package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"path/filepath"
    "regexp"
	"flag"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
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
	configPath := flag.String("config", "", "path to config file (e.g., config.json)")
	flag.Parse()

	if *configPath == "" {
		log.Fatal("Please provide a config file path using --config flag")
	}

	config, err := readConfigFile(*configPath)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	myFigure := figure.NewFigure(config.Message, "", true)
	myFigure.Print()

	gi, _ := GetInfo()
	gi.VarDump()

	// Print memory usage
	memInfo := MemInfo{}
	memInfo.Update()
	goMemInfo := ConvertToStruct(&memInfo)
	memTotalGB := float64(goMemInfo.MemTotal) / (1024 * 1024 * 1024)
	memAvailableGB := float64(goMemInfo.MemAvailable) / (1024 * 1024 * 1024)
	color.Green(fmt.Sprintf("%-20s: %.2f GB", "Memory Info", memTotalGB))
	color.Green(fmt.Sprintf("%-20s: %.2f GB", "Memory Available", memAvailableGB))
	uptime, err := getUptime()
	if err != nil {
		log.Printf("Error getting uptime: %v", err)
	}
	color.Green(fmt.Sprintf("%-20s: %.2f hours", "Up Time", uptime.Hours()))
}
