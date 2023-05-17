package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type dirFunc func() (string, error)
type dirCheckFunc func(string) error

func Replace4Store(path string) (string, error) {
	return _replace4store(
		path,
		func() (string, error) { return os.UserHomeDir() },
		func() (string, error) { return os.Getwd() },
		func(path string) error {
			if !Exists(path) {
				return fmt.Errorf("no dir found for path: [%s]\n", path)
			}
			return nil
		},
	)
}

func _replace4store(path string, hdF, wdF dirFunc, chF dirCheckFunc) (string, error) {
	hd, err := hdF()
	if err != nil {
		return "", err
	}
	wd, err := wdF()
	if err != nil {
		return "", err
	}
	if path == "." {
		return replaceHome(wd, hd)
	}
	if path == "$HOME" || path == "${HOME}" {
		return path, nil
	}
	if isRelPath := !strings.HasPrefix(path, "/"); isRelPath {
		joinedPath := filepath.Join(wd, path)
		if err := chF(joinedPath); err != nil {
			return "", err
		}
		return replaceHome(joinedPath, hd)
	}
	return replaceHome(path, hd)
}

func replaceHome(path, homeDir string) (string, error) {
	if isLeafOfHD := strings.HasPrefix(path, homeDir); isLeafOfHD {
		shortenedPath := filepath.Join("${HOME}", strings.TrimPrefix(path, homeDir))
		return shortenedPath, nil
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
