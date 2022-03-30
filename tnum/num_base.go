package tnum

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
)

//取大于min小于 max 的 随机整数
func Randint(min int64, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return n.Int64() + min
}

func Random() float64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(99749965075508026))

	return float64(n.Int64()) / float64(math.Pow10(17))
}

//返回 百分比  例 0.051  返回 5.1
func Percent(p float64) float64 {
	return math.Round(p*10000) / 100
}

//四四舍五入 保留指定长度
func Floor(num float64, l int) float64 {
	s2 := strconv.FormatFloat(num, 'f', l, 64)

	f64, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return 0
	}
	return f64
}
