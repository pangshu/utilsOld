package File

import "path/filepath"

// Glob 寻找与模式匹配的文件路径.
func (*File)Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}
