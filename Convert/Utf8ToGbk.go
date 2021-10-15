package Convert

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// Utf8ToGbk UTF-8转GBK编码.
func (*Convert)Utf8ToGbk(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	return string(d), e
}
