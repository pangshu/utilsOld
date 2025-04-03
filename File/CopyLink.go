package File

import (
	"os"
	"path/filepath"
)

// CopyLink 拷贝链接.
func (toolFile *File)CopyLink(source string, dest string) error {
	if source == dest {
		return nil
	}

	source, err := os.Readlink(source)
	if err != nil {
		return err
	}

	_, err = os.Lstat(dest)
	if err == nil {
		_ = os.Remove(dest)
	}

	//创建目录
	destDir := filepath.Dir(dest)
	if !toolFile.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return err
		}
	}

	return os.Symlink(source, dest)
}
