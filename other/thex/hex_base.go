package thex

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func GetMd5(bytes []byte) string {
	has := md5.Sum(bytes)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

//base64编码
func EncodeBase64(str string) string {
	input := []byte(str)
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString
}

//base64解码
func DecodeBase64(str string) string {
	decodeString, _ := base64.StdEncoding.DecodeString(str)
	return string(decodeString)
}
