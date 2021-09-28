package tfile

import (
	"fmt"
	"testing"
)

func Test_PathGetFileName(t *testing.T) {

	fmt.Println(PathGetFileName("/device/sdk/CMakeLists.txt", true))
	fmt.Println(PathGetFileName("/device/sdk/CMakeLists.txt", false))

	fmt.Println(PathGetFileName("https://raw.githubusercontent.com/Aquietorange/man2v/master/test/NetPenetrate.sh", true))
	fmt.Println(PathGetFileName("https://raw.githubusercontent.com/Aquietorange/man2v/master/test/NetPenetrate.sh", false))
}
