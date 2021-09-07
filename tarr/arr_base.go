package tarr

//查找指定 文本元素 索引
func Strsearch(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
