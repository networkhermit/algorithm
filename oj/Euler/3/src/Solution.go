package main

import (
    "fmt"
    "io"
    "math"
)

func process(n int64) int64 {
    if n < 2 {
        return 1
    }

    for (n & 1) == 0 {
        n >>= 1
    }

    if n == 1 {
        return 2
    }

    for i, bound := int64(3), int64(math.Sqrt(float64(n))) + 1; i < bound; i++ {
        for n % i == 0 {
            if n == i {
                return n
            }
            n /= i
        }
    }

    return n
}

func main() {
    var testCases int

    var n int64

    for {
        _, err := fmt.Scan(&testCases)
        if err == io.EOF {
            break
        }
        for i := 0; i < testCases; i++ {
            fmt.Scan(&n)
            fmt.Println(process(n))
        }
    }
}
