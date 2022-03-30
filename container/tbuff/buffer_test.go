package tbuff

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_NewBuffer(T *testing.T) {
	bins, _ := ioutil.ReadFile("F:\\vuex\\go_demo\\English\\workbox\\test.zpk")
	Buf := NewBuffer(bins)
	Buf.Setoffset(128)
	fmt.Println(Buf.ReadHex(11))
	ss := Buf.ReadString(66)
	fmt.Println(ss)
	Buf.SetoffsetHex("8000000000000000")
	Buf.Readbytes(8)             //00 CD D9 57 2B FF 90 21 84
	fmt.Println(Buf.ReadInt32()) //D8 02  00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //A9 86 01 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //A0 AF 00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //51 1D 00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //54 10 00 00
	fmt.Println("-----")
	Buf.close()

	Buf, _ = NewBufferByFile("F:\\vuex\\go_demo\\English\\workbox\\test.zpk")
	Buf.SetoffsetHex("8000000000000000")
	Buf.Readbytes(8)             //00 CD D9 57 2B FF 90 21 84
	fmt.Println(Buf.ReadInt32()) //D8 02  00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //A9 86 01 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //A0 AF 00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //51 1D 00 00
	Buf.Readbytes(44)
	fmt.Println(Buf.ReadInt32()) //54 10 00 00
	fmt.Println("-----")
	Buf.close()

	//05 00 00 00  文件数
	//46 67 02 各 文件长度和信息 的起始位置
}
