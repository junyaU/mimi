package utils

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// GetModuleName executes "go list -m" command and returns the module name.
// It trims whitespace from the output and converts it to a string.
func GetModuleName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// IsMatchedPackage checks if the package name contains the module path from the file path.
// It returns false if there is an error getting the module name or if the file path is empty.
func IsMatchedPackage(filePath, packageName string) bool {
	if filePath == "" {
		return false
	}

	moduleName, err := GetModuleName()
	if err != nil {
		return false
	}

	pkgPath := strings.ReplaceAll(filepath.Clean(filePath), "../", "")

	modulePath := filepath.Join(moduleName, pkgPath)

	return strings.Contains(packageName, modulePath)
}

// Contains checks if the slice contains the target string.
// It returns true if the target string is found in the slice, false otherwise.
func Contains(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}
