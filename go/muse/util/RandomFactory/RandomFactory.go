package RandomFactory

import (
    "math/rand"
    "time"
)

func Seed() {
    rand.Seed(time.Now().UnixNano())
}

func GenIntN(min int, max int) int {
    return min + rand.Intn(max - min + 1)
}

func GenInt() int {
    return GenIntN(0, 2147483647)
}

func GenEven() int {
    return GenInt() >> 1 << 1
}

func GenOdd() int {
    return GenInt() | 1
}
