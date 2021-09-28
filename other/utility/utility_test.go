package utility

import (
	"fmt"
	"math"
	"testing"
)

func Test_SortFast(T *testing.T) {
	fmt.Println(SortFastInt([]int{
		3, 4, 5, 6, 46, 357, 3, 537, 357, 2, 24, 753, 883, 68997, 35, 246, 246, 24, 735, 848, 25, 8, 3524, 5, 735, 7, 357, 8, 46, 84, 68, 34, 57, 57, 3, 68, 76, 8, 458, 46, 845, 8, 5, 36,
	}))
	var aa float64 = (65 - 60) / float64(30)
	fmt.Println(aa)
	var offd int64 = 66

	offd = offd + int64(math.Ceil(1.0/6.0))
	fmt.Println(offd)
}
