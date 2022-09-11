package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetGoModuleVersion() (string, error) {
	version := os.Getenv("GOVERSION")
	if version != "" {
		return formatGoModuleVersion(version)
	}
	cmd := exec.Command("go", "version")
	b, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return formatGoModuleVersion(strings.Split(string(b), " ")[2])
}

func formatGoModuleVersion(v string) (string, error) {
	s := strings.Split(strings.ReplaceAll(v, "go", ""), ".")
	if len(s) < 2 {
		return "", errors.New("go version is not available")
	}
	return fmt.Sprintf("%s.%s", s[0], s[1]), nil // format to `1.23`
}
