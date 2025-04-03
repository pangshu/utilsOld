package File

import "os"

// 获取项目路径
func (*File) ProjectPath() string{
    var projectPath string
    projectPath, _ = os.Getwd()
    return projectPath
}
