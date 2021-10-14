package File

import (
	"path/filepath"
	"strings"
)

// Tree 获取目录的文件树列表.
// ftype为枚举(FILE_TREE_ALL、FILE_TREE_DIR、FILE_TREE_FILE);
// recursive为是否递归;
// filters为一个或多个文件过滤器函数,FileFilter类型.

type (
	FileFilter func(string) bool
)

func (toolFile *File)Tree(filePath string, fileType int, recursive bool, filters ...FileFilter) []string {
	var trees []string

	if toolFile.IsFile(filePath, 0) {
		if fileType != TreeDir {
			trees = append(trees, filePath)
		}
		return trees
	}

	filePath = strings.TrimRight(filePath, "/")
	files, err := filepath.Glob(filepath.Join(filePath, "*"))
	if err != nil || len(files) == 0 {
		return trees
	}

	for _, file := range files {
		//文件过滤
		chk := true
		if len(filters) > 0 {
			for _, filter := range filters {
				chk = filter(file)
				if !chk {
					break
				}
			}
		}
		if !chk {
			continue
		}

		if toolFile.IsDir(file) {
			if fileType != TreeFile {
				trees = append(trees, file)
			}

			if recursive {
				subs := toolFile.Tree(file, fileType, recursive)
				trees = append(trees, subs...)
			}
		} else if fileType != TreeDir {
			trees = append(trees, file)
		}
	}

	return trees
}
