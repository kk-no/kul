package internal

import (
	"os"
	"os/exec"
	"strings"
)

func GetGoModuleVersion() (string, error) {
	version := os.Getenv("GOVERSION")
	if version != "" {
		return formatGoModuleVersion(version), nil
	}
	cmd := exec.Command("go", "version")
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return formatGoModuleVersion(strings.Split(string(b), " ")[2]), nil
}

func formatGoModuleVersion(v string) string {
	return strings.ReplaceAll(v, "go", "")
}
