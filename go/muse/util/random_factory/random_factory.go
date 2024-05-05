package random_factory

import (
	"math/rand/v2"
)

func GenIntN(min, max int) int {
	return min + rand.N(max-min+1)
}

func GenInt() int {
	return GenIntN(1, 2_147_483_647)
}

func GenEven() int {
	return GenInt() >> 1 << 1
}

func GenOdd() int {
	return GenInt() | 1
}
