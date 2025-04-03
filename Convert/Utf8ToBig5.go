package Convert

import (
	"bytes"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// Utf8ToBig5 UTF-8转BIG5编码.
func (*Convert)Utf8ToBig5(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	return string(d), e
}
