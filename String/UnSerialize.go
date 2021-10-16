package String

import (
	"bytes"
	"encoding/gob"
)

// UnSerialize 对字符串进行反序列化;其中registers注册对象,其类型必须和Serialize的一致.
func (*String)UnSerialize(data []byte, registers ...interface{}) (val interface{}, err error) {
	for _, v := range registers {
		gob.Register(v)
	}

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(&val)
	return
}
