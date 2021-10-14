package File

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"strings"
)

// UnZip 解压zip文件.srcZip为zip文件路径,dstDir为解压目录.
func (toolFile *File)UnZip(srcZip, dstDir string) (bool, error) {
	reader, err := zip.OpenReader(srcZip)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = reader.Close()
	}()

	dstDir = strings.TrimRight(toolFile.AbsPath(dstDir), "/\\")
	if !toolFile.IsExist(dstDir) {
		err := toolFile.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	// 迭代压缩文件中的文件
	for _, f := range reader.File {
		// Create diretory before create file
		newPath := dstDir + "/" + strings.TrimLeft(f.Name, "/\\")
		parentDir := path.Dir(newPath)
		if !toolFile.IsExist(parentDir) {
			err := os.MkdirAll(parentDir, os.ModePerm)
			if err != nil {
				return false, err
			}
		}

		if !f.FileInfo().IsDir() {
			if fcreate, err := os.Create(newPath); err == nil {
				if rc, err := f.Open(); err == nil {
					_, _ = io.Copy(fcreate, rc)
					_ = rc.Close() //不要用defer来关闭，如果文件太多的话，会报too many open files 的错误
					_ = fcreate.Close()
				} else {
					_ = fcreate.Close()
					return false, err
				}
			} else {
				return false, err
			}
		}
	}

	return true, nil
}

