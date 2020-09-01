package file

import (
	"os"
	"path/filepath"
	"strings"
)

// AbsolutePath 获取文件绝对路径
func AbsolutePath(filePath string) string {
	path, _ := filepath.Abs(filePath)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}

// Exists 判断文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
