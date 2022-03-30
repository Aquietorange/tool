package tfile

import (
	"archive/zip"
	"io"
	"io/ioutil"
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
	if strings.Contains(pathf, "\\") {
		ss := strings.Split(pathf, "\\")
		pathf = ss[len(ss)-1]
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

// IsFile checks whether the path is a file,
// it returns false when it's a directory or does not exist.
// 为目录 或不存在时 返回假
func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil { //不存在
		return false
	}
	return !fi.IsDir()
}

// IsDir checks whether the path is a dir,
// it returns false when it's a directory or does not exist.
// 为文件 或不存在时 返回假
func IsDir(f string) bool {
	fi, e := os.Stat(f)
	if e != nil { //不存在
		return false
	}
	return fi.IsDir()
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

//获取指定目录下的所有子目录，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDirPaths(dirPth string) (paths []string, err error) {
	paths = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			paths = append(paths, dirPth+PthSep+fi.Name())
		}

	}
	return paths, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})
	return files, err
}

/* //解压
func DeZIPCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	os.MkdirAll(dest, 0666)
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name

		err = os.MkdirAll(path.Dir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
} */

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func ZipCompress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := zipcompress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func zipcompress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = zipcompress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeZIPCompress(zipFile, dest_path string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	DeCompress_to_dest := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest_path + f.Name
		if err = os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		return err
	}
	for _, file := range reader.File {
		err := DeCompress_to_dest(file)
		if err != nil {
			return err
		}
	}
	return nil
}
