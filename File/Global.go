package File

type File byte

const (
	// FILE_COVER_ALLOW 文件覆盖,允许
	CoverAllow = 1
	// FILE_COVER_IGNORE 文件覆盖,忽略
	CoverIgnore = 0
	// FILE_COVER_DENY 文件覆盖,禁止
	CoverDeny = -1

	// FILE_TREE_ALL 文件树,查找所有(包括目录和文件)
	TreeAll = 3
	// FILE_TREE_DIR 文件树,仅查找目录
	TreeDir = 2
	// FILE_TREE_FILE 文件树,仅查找文件
	TreeFile = 1
)