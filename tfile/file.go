package tfile

import (
	"os"
	"path/filepath"
	"strings"
)

//判断 文件或目录是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//程序运行目录 失败返回空
func GetCurrentDirectory() (string, error) {
	//返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1), nil
}
