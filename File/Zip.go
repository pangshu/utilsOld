package File

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Zip 将文件目录进行zip打包.fpaths为文件或目录的路径.
func (toolFile *File)Zip(dst string, fpaths ...string) (bool, error) {
	dst = toolFile.AbsPath(dst)
	dstDir := toolFile.Dirname(dst)
	if !toolFile.IsExist(dstDir) {
		_ = toolFile.Mkdir(dstDir, os.ModePerm)
	}

	fzip, err := os.Create(dst)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fzip.Close()
	}()

	if len(fpaths) == 0 {
		return false, errors.New("No input files.")
	}

	var allfiles, files []string
	var fpath string
	for _, fpath = range fpaths {
		fpath = strings.Trim(fpath, " \t\n\r\v\f\x00　")
		if toolFile.IsDir(fpath) {
			files = toolFile.Tree(fpath, 3, true)
			if len(files) != 0 {
				allfiles = append(allfiles, files...)
			}
		} else if fpath != "" {
			allfiles = append(allfiles, fpath)
		}
	}

	if len(allfiles) == 0 {
		return false, errors.New("No exist files.")
	}

	zipw := zip.NewWriter(fzip)
	defer func() {
		_ = zipw.Close()
	}()

	keys := make(map[string]bool)
	for _, fpath = range allfiles {
		if _, ok := keys[fpath]; ok || toolFile.AbsPath(fpath) == dst {
			continue
		}

		fileToZip, err := os.Open(fpath)
		if err != nil {
			return false, fmt.Errorf("Failed to open %s: %s", fpath, err)
		}
		defer func() {
			_ = fileToZip.Close()
		}()

		wr, _ := zipw.Create(fpath)
		keys[fpath] = true
		if _, err := io.Copy(wr, fileToZip); err != nil {
			return false, fmt.Errorf("Failed to write %s to zip: %s", fpath, err)
		}
	}

	return true, nil
}

