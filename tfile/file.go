package tfile

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Aquietorange/tool/tstr"
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

//从路径或url 取文件名 ,needSuffixes 返回值是否 包含后辍
func PathGetFileName(pathf string, needSuffixes bool) string {

	if strings.Contains(pathf, "//") {
		pathf = tstr.Substr_str(pathf, "//", -1)
	}

	//filename := "device/sdk/CMakeLists.txt"
	filenameall := path.Base(pathf)
	filesuffix := path.Ext(pathf) //后辍
	fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]

	if needSuffixes {
		return filenameall
	} else {
		return fileprefix
	}
}

//取 文件路径或url  的 文件后辍
func GetFileSuffixes(file string) string {
	if strings.Contains(file, "//") {
		file = tstr.Substr_str(file, "//", -1)
	}
	//filename := "device/sdk/CMakeLists.txt"
	//filenameall := path.Base(file)
	filesuffix := path.Ext(file) //后辍
	//fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
	return filesuffix
}
