package String

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// ShaX 计算字符串的 shaX 散列值,x为1/256/512 .
func (*String)ShaX(str string, x ...int) string {
	xNum := 1
	if len(x) > 0 {
		if x[0] == 1 || x[0] == 256 || x[0] == 512 {
			xNum = x[0]
		}
	}
	data := []byte(str)
	var h hash.Hash
	switch xNum {
	case 1:
		h = sha1.New()
		break
	case 256:
		h = sha256.New()
		break
	case 512:
		h = sha512.New()
		break
	default:
		h = sha1.New()
		//panic("[shaXByte] x must be in [1, 256, 512]")
	}

	h.Write(data)

	hBytes := h.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return string(res)
}

