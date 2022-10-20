package pkg

import (
	"path"
	"strings"
)

// GetFilePath 去除后缀
func GetFilePath(filePath string) string {
	ext := path.Ext(filePath)
	return strings.TrimSuffix(filePath, ext)
}
