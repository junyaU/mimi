package utils

import (
	"os/exec"
	"path/filepath"
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

func Contains(slice []string, target string) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}
