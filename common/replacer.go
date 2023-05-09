package common

import (
	"os"
	"path/filepath"
	"strings"
)

func Replace4Store(path string) (string, error) {
	if path == "." {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		hd, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = wd
		if strings.HasPrefix(path, hd) {
			path = filepath.Join("${HOME}", strings.TrimPrefix(path, hd))
		}
		return path, nil
	}

	return path, nil
}

func Replace4Get(path string) (string, error) {
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
