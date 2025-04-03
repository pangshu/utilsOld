package File

import (
	"fmt"
	"os"
	"path/filepath"
)

// CopyDir 拷贝源目录到目标目录,cover为枚举(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (toolFile *File)CopyDir(source string, dest string, coverType ...int) (int64, error) {
	var total, nBytes int64
	var err error
	cover := CoverAllow
	if len(coverType) > 0 {
		cover = coverType[0]
	}
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return 0, err
	} else if !sourceInfo.IsDir() {
		return 0, fmt.Errorf("%s is not a directory", source)
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return 0, err
	}

	directory, _ := os.Open(source)
	defer func() {
		_ = directory.Close()
	}()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return 0, err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			nBytes, err = toolFile.CopyDir(srcFilePath, destFilePath, cover)
		} else {
			destFileInfo, err := os.Stat(destFilePath)
			if err == nil {
				if cover != CoverAllow || os.SameFile(obj, destFileInfo) {
					continue
				}
			}

			if obj.Mode()&os.ModeSymlink != 0 {
				// a link
				_ = toolFile.CopyLink(srcFilePath, destFilePath)
			} else {
				nBytes, err = toolFile.CopyFile(srcFilePath, destFilePath, cover)
			}
		}

		if err == nil {
			total += nBytes
		}
	}

	return total, err
}
