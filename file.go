package utils

import (
	"os"
	"path/filepath"
)

// current path
func CurrentPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}
