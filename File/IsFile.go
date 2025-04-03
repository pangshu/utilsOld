package File

import "os"

type ParamFileType uint8
const (
	// FILE_TYPE_ANY 文件类型-任意
	FILE_TYPE_ANY = 0
	// FILE_TYPE_LINK 文件类型-链接文件
	FILE_TYPE_LINK = 1
	// FILE_TYPE_REGULAR 文件类型-常规文件(不包括链接)
	FILE_TYPE_REGULAR = 2
	// FILE_TYPE_COMMON 文件类型-普通文件(包括常规和链接)
	FILE_TYPE_COMMON = 3
)

// IsFile 判断文件是否为某类型的文件，且存在
func (toolFile *File)IsFile(filePath string, fileType ...ParamFileType) (res bool) {
	isBool := toolFile.IsExist(filePath)
	if true == isBool {
		var t ParamFileType = FILE_TYPE_ANY
		if len(fileType) > 0 {
			t = fileType[0]
		}

		var f os.FileInfo
		var e error
		var musLink, musRegular bool

		if t == FILE_TYPE_LINK {
			musLink = true
		} else if t == FILE_TYPE_REGULAR {
			musRegular = true
		} else if t == FILE_TYPE_COMMON {
			musLink = true
			musRegular = true
		}
		if (!musLink && !musRegular) || musRegular {
			f, e := os.Stat(filePath)
			if musRegular {
				res = (e == nil) && f.Mode().IsRegular()
			} else {
				res = (e == nil) && !f.IsDir()
			}
		}

		if !res && musLink {
			f, e = os.Lstat(filePath)
			res = (e == nil) && (f.Mode()&os.ModeSymlink == os.ModeSymlink)
		}
		return
	} else {
		return false
	}
}
