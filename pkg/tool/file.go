package tool

import (
	"os"
	"path/filepath"
	"strings"
)

// GetAppDirectory 获取程序运行路径
func GetAppDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1), err
}
