package File

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// TarGz 打包压缩tar.gz;src为源文件或目录,dstTar为打包的路径名,ignorePatterns为要忽略的文件正则.
func (toolFile *File)TarGz(src string, dstTar string, ignorePatterns ...string) (bool, error) {
	//过滤器,检查要忽略的文件
	var filter = func(file string) bool {
		res := true
		for _, pattern := range ignorePatterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				continue
			}
			chk := re.MatchString(file)
			if chk {
				res = false
				break
			}
		}
		return res
	}

	src = toolFile.AbsPath(src)
	dstTar = toolFile.AbsPath(dstTar)

	dstDir := toolFile.Dirname(dstTar)
	if !toolFile.IsExist(dstDir) {
		_ = toolFile.Mkdir(dstDir, os.ModePerm)
	}

	files := toolFile.Tree(src, 3, true, filter)
	if len(files) == 0 {
		return false, fmt.Errorf("src no files to tar.gz")
	}

	// dest file write
	fw, err := os.Create(dstTar)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fw.Close()
	}()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer func() {
		_ = gw.Close()
	}()

	// tar write
	tw := tar.NewWriter(gw)
	defer func() {
		_ = tw.Close()
	}()

	parentDir := filepath.Dir(src)
	for _, file := range files {
		if file == dstTar {
			continue
		}
		fi, err := os.Stat(file)
		if err != nil {
			continue
		}
		newName := strings.Replace(file, parentDir, "", -1)

		// Create tar header
		hdr := new(tar.Header)
		hdr.Format = tar.FormatGNU

		if fi.IsDir() {
			// if last character of header name is '/' it also can be directory
			// but if you don't set Typeflag, error will occur when you untargz
			hdr.Name = newName + "/"
			hdr.Typeflag = tar.TypeDir
			hdr.Size = 0
			//hdr.Mode = 0755 | c_ISDIR
			hdr.Mode = int64(fi.Mode())
			hdr.ModTime = fi.ModTime()

			// Write hander
			err := tw.WriteHeader(hdr)
			if err != nil {
				return false, fmt.Errorf("DirErr: %s file:%s\n", err.Error(), file)
			}
		} else {
			// File reader
			fr, err := os.Open(file)
			if err != nil {
				return false, fmt.Errorf("OpenErr: %s file:%s\n", err.Error(), file)
			}
			defer func() {
				_ = fr.Close()
			}()

			hdr.Name = newName
			hdr.Size = fi.Size()
			hdr.Mode = int64(fi.Mode())
			hdr.ModTime = fi.ModTime()

			// Write hander
			err = tw.WriteHeader(hdr)
			if err != nil {
				return false, fmt.Errorf("FileErr: %s file:%s\n", err.Error(), file)
			}

			// Write file data
			_, err = io.Copy(tw, fr)
			if err != nil {
				return false, fmt.Errorf("CopyErr: %s file:%s\n", err.Error(), file)
			}
			_ = fr.Close()
		}
	}

	return true, nil
}

