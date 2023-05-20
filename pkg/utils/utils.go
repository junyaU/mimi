package utils

import (
	"os/exec"
	"strings"
)

func GetModuleName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func Contains(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func Max(slice []int) int {
	max := 0
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
