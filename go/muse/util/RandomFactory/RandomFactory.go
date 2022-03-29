package RandomFactory

import (
	"crypto/rand"
	"math/big"
)

func GenIntN(min int, max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return min + int(n.Int64())
}

func GenInt() int {
	return GenIntN(1, 2147483647)
}

func GenEven() int {
	return GenInt() >> 1 << 1
}

func GenOdd() int {
	return GenInt() | 1
}
