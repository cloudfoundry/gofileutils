package fileutils

import (
	"path/filepath"
	"strings"
)

// see MAX_PATH limitation in https://msdn.microsoft.com/en-us/library/windows/desktop/aa364946%28v=vs.85%29.aspx
const (
	extPathLenPrefix = `\\?\`
)

func AbsPath(path string) (string, error) {
	p, err := filepath.Abs(path)
	if err != nil {
		return p, err
	}
				
	if !strings.HasPrefix(p, extPathLenPrefix) {
		p = extPathLenPrefix + p
	}

	return p, nil
}
