package thex

import (
	"fmt"
	"testing"
)

func Test_EncodeBase64(t *testing.T) {
	aa := EncodeBase64("aaaaaaaaaadd")
	fmt.Println(aa)
	fmt.Println(DecodeBase64(aa))
}
