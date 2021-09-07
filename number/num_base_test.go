package tnum

import (
	"fmt"
	"testing"
)

func Test_Randint(t *testing.T) {
	rint := Randint(1, 10000)
	//uuid, _ = ParseString(uuid.String())
	fmt.Println(rint)
}
