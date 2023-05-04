package common

import (
	"os"
	"strings"
)

func Replace(path string) (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(path, "${HOME}") {
		path = strings.ReplaceAll(path, "${HOME}", dir)
	}
	if strings.HasPrefix(path, "$HOME") {
		path = strings.ReplaceAll(path, "$HOME", dir)
	}
	if strings.HasPrefix(path, "~") {
		path = strings.ReplaceAll(path, "~", dir)
	}
	return path, nil
}
