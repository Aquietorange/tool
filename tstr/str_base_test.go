package tstr

import (
	"fmt"
	"testing"
)

func Test_Substr(t *testing.T) {
	fmt.Println(Substr("abcd【名】 （Roller）（英、德、俄、匈、罗、捷、瑞典）罗勒，（法）罗莱（人名）", 0, 16))
}
