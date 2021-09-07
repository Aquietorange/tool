package tstr

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
