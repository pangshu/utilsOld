package String

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
)

// IsRsaPublicKey 检查字符串是否RSA的公钥,keylen为密钥长度.
func IsRsaPublicKey(str string, keylen int) bool {
	bb := bytes.NewBufferString(str)
	pemBytes, _ := ioutil.ReadAll(bb)

	// 获取公钥
	block, _ := pem.Decode(pemBytes)
	if block != nil && block.Type != "PUBLIC KEY" {
		return false
	}
	var der []byte
	var err error

	if block != nil {
		der = block.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return false
		}
	}

	key, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return false
	}
	pubkey, ok := key.(*rsa.PublicKey)
	if !ok {
		return false
	}
	bitlen := len(pubkey.N.Bytes()) * 8
	return bitlen == int(keylen)
}
