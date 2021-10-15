package Convert

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// GbkToUtf8 GBK转UTF-8编码.
func (*Convert)GbkToUtf8(str string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	return string(d), e
}
