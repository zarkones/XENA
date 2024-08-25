package wordlists

var WebPathFileUpload = func() []string {
	level1 := []string{
		"upload",
		"file-upload",
		"file_upload",
		"upload-file",
		"upload_file",
		"uploadFile",
		"fileUpload",
		"file/upload",
		"upload/file",
	}
	level2 := []string{
		"v1/",
		"v2/",
		"v3/",
		"api/v1/",
		"api/v2/",
		"api/v3/",
		"v1/api/",
		"v2/api/",
		"v3/api/",
	}
	paths := level1
	for _, l1 := range level1 {
		for _, l2 := range level2 {
			paths = append(paths, l2+l1)
		}
	}
	return paths
}()
