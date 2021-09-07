package tnum

import (
	"crypto/rand"
	"math/big"
)

//取大于min小于 max 的 随机整数
func Randint(min int64, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max))
	return n.Int64() + min
}
