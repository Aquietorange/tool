package tfile

import (
	"fmt"
	"os"
	"testing"
)

func Test_PathGetFileName(t *testing.T) {
	fmt.Println(PathGetFileName("F:\\教程资料\\学英语\\百词斩\\410\\zp_10087_410_0_20210525151531.zpk", true))
	fmt.Println(PathGetFileName("/device/sdk/CMakeLists.txt", true))
	fmt.Println(PathGetFileName("/device/sdk/CMakeLists.txt", false))

	fmt.Println(PathGetFileName("https://raw.githubusercontent.com/Aquietorange/man2v/master/test/NetPenetrate.sh", true))
	fmt.Println(PathGetFileName("https://raw.githubusercontent.com/Aquietorange/man2v/master/test/NetPenetrate.sh", false))
}

func Test_DeZIPCompress(t *testing.T) {
	DeZIPCompress("C:\\Users\\hujun\\Documents\\netpe_win64.zip", "C:\\Users\\hujun\\Documents\\netpe_win64\\")

}

func Test_ZipCompress(t *testing.T) {
	var fs []*os.File
	f, _ := os.Open("C:\\Users\\hujun\\Documents\\netpe_win64\\")
	fs = append(fs, f)
	ZipCompress(fs, "C:\\Users\\hujun\\Documents\\netpe_win64.zip")
}
