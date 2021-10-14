package File

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// DelDir 删除目录;delRoot为true时连该目录一起删除;为false时只清空该目录.
func (toolFile *File)DelDir(dir string, delRoot ...bool) error {
	delDir := true
	if len(delRoot) > 0 {
		delDir = delRoot[0]
	}
	realPath := toolFile.AbsPath(dir)
	if !toolFile.IsDir(realPath) {
		return fmt.Errorf("Dir %s not exists", dir)
	}

	names, err := ioutil.ReadDir(realPath)
	if err != nil {
		return err
	}

	for _, v := range names {
		file := path.Join([]string{realPath, v.Name()}...)
		err = os.RemoveAll(file)
	}

	//删除根节点(指定的目录)
	if delDir {
		err = os.RemoveAll(realPath)
	}

	return err
}
