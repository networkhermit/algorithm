package RandomFactory

import (
    "math/rand"
    "time"
)

func Launch() {
    rand.Seed(time.Now().UnixNano())
}

func IntegerN(min int, max int) int {
    return min + rand.Intn(max - min)
}

func GenerateInteger() int {
    return IntegerN(0, 2147483647)
}

func GenerateEven() int {
    return GenerateInteger() >> 1 << 1
}

func GenerateOdd() int {
    return GenerateInteger() | 1
}
