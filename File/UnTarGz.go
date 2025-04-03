package File

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// UnTarGz 将tar.gz文件解压缩;srcTar为压缩包,dstDir为解压目录.
func (toolFile *File)UnTarGz(srcTar, dstDir string) (bool, error) {
	fr, err := os.Open(srcTar)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fr.Close()
	}()

	dstDir = strings.TrimRight(toolFile.AbsPath(dstDir), "/\\")
	if !toolFile.IsExist(dstDir) {
		err := toolFile.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	// Gzip reader
	gr, err := gzip.NewReader(fr)

	// Tar reader
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		} else if err != nil {
			return false, err
		}

		// Create diretory before create file
		newPath := dstDir + "/" + strings.TrimLeft(hdr.Name, "/\\")
		parentDir := path.Dir(newPath)
		if !toolFile.IsExist(parentDir) {
			_ = os.MkdirAll(parentDir, os.ModePerm)
		}

		if hdr.Typeflag != tar.TypeDir {
			// Write data to file
			fw, err := os.Create(newPath)
			if err != nil {
				return false, fmt.Errorf("CreateErr: %s file:%s\n", err.Error(), newPath)
			}

			_, err = io.Copy(fw, tr)
			if err != nil {
				return false, fmt.Errorf("CopyErr: %s file:%s\n", err.Error(), newPath)
			}

			_ = fw.Close()
		}
	}

	return true, nil
}
