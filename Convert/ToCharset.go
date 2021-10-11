package Convert

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// Charset 强制将变量转换为整型.
func (*Convert)Charset(src string, srcCharset string, dstCharset string) (dst string,err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src

	// Converting <src> to UTF-8.
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, tmpErr := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if tmpErr != nil {
				return "", fmt.Errorf("%s to utf8 failed. %v", srcCharset, err)
			}
			src = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupport srcCharset: %s", srcCharset))
		}
	}

	// Do the converting from UTF-8 to <dstCharset>.
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, tmpErr := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if tmpErr != nil {
				return "", fmt.Errorf("utf to %s failed. %v", dstCharset, err)
			}
			dst = string(tmp)
		} else {
			return dst, errors.New(fmt.Sprintf("unsupport dstCharset: %s", dstCharset))
		}
	} else {
		dst = src
	}
	return dst, nil
}

func getEncoding(charset string) encoding.Encoding {
	charsetAlias := map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}

	if e, err := ianaindex.MIB.Encoding(charset); err == nil && e != nil {
		return e
	}
	return nil
}