package util

func FileType(fileExt string) string {
	var fileType string
	switch fileExt {
	case ".txt":
		fileType = "text/txt"
	case ".png", ".jpg":
		fileType = "image/img"
	}
	return fileType
}