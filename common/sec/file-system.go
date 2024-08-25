package sec

func SearchFileSystem(term string) (filePaths []string, err error) {
	return filePaths, nil // TODO
}

// File System Entry.
type FSEntry struct {
	Name  bool `json:"name"`
	IsDir bool `json:"isDir"`
}

func BrowsePath(path string) (fileEntries []FSEntry, err error) {
	return fileEntries, nil // TODO
}
