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
func Test_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		rint := Random()
		fmt.Println(rint)
	}

}
