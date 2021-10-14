package File

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 获取文件md5值,length指定结果长度32/16.
func (*File)Md5(filePath string, length ...int) (string, error) {
	var num int
	if len(length) == 0 {
		num = 32
	} else {
		if length[0] != 16 && length[0] != 32 {
			num = 32
		} else {
			num = length[0]
		}
	}

	var res string
	f, err := os.Open(filePath)
	if err != nil {
		return res, err
	}
	defer func() {
		_ = f.Close()
	}()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return res, err
	}

	hashInBytes := hash.Sum(nil)
	if num == 16 || num == 32 {
		dst := make([]byte, hex.EncodedLen(len(hashInBytes)))
		hex.Encode(dst, hashInBytes)
		if num == 16 {
			res = string(dst[8:24])
		} else {
			res = string(dst)
		}
	} else {
		res = hex.EncodeToString(hashInBytes)
	}

	return res, nil
}
