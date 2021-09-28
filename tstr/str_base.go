package tstr

import (
	"crypto/rand"
	"math/big"
	"strings"
)

//截取文本 str =开始位置,len=截取长度, -1=截取到尾部
func Substr(s string, str, lena int) string {
	ss := []byte(s)
	if str > len(ss) {
		return ""
	}
	if lena == -1 {
		return string(ss[str:])
	} else {
		return string(ss[str : str+lena])
	}
}

//截取文本 str =开始截取的文本(不含),len=截取长度, -1=截取到尾部
func Substr_str(s string, strs string, lena int) string {
	str := strings.Index(s, strs)
	if str == -1 {
		return ""
	}

	ss := []byte(s)
	if str > len(ss) {
		return ""
	}
	if lena == -1 {
		return string(ss[str:])
	} else {
		return string(ss[str : str+lena])
	}
}

//取随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b, _ := rand.Int(rand.Reader, big.NewInt(26))
		bytes[i] = byte(b.Int64() + 65)
	}
	return string(bytes)
}
