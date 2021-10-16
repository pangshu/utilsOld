package String

import (
	"bytes"
	"encoding/gob"
)

// Serialize 对变量进行序列化.
func (*String)Serialize(val interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	gob.Register(val)

	err := enc.Encode(&val)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}