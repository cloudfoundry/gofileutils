//

// +build !windows

package fileutils

import (
	"path/filepath"
)

func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}
