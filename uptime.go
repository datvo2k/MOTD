package main

import (
	"os"
	"strconv"
	"strings"
	"time"
	"errors"
)

func getUptime() (time.Duration, error) {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return 0, err
	}

	// Split the content to get the uptime value (first number)
	fields := strings.Fields(string(content))
	if len(fields) < 1 {
		return 0, errors.New("invalid /proc/uptime format")
	}

	// Parse the uptime value (it's in seconds)
	uptime, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	return time.Duration(uptime * float64(time.Second)), nil
}