package machine

import (
	"os"
	"path/filepath"
)

func WalkDir(path string) (paths []string, err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return paths, err
	}

	for _, entry := range entries {
		currPath := filepath.Join(path, entry.Name())
		if entry.IsDir() {
			subEntries, _ := WalkDir(currPath)
			paths = append(paths, subEntries...)
			continue
		}
		paths = append(paths, currPath)
	}

	return paths, err
}
