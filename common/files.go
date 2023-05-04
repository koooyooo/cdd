package common

import (
	"github.com/koooyooo/cdd/common/constant"
	"os"
	"path/filepath"
)

// Exists returns true if the file or dir exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CDDPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, constant.CDDFilePath), nil
}
