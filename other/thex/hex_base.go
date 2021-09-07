package thex

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(bytes []byte) string {
	has := md5.Sum(bytes)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
