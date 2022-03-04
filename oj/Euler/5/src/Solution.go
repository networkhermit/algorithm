package main

import (
    "fmt"
    "io"
    "math"
)

func process(n int) int64 {
    if n < 2 {
        return 1
    }

    size := (n + 1) >> 1

    sievedFactors := make([]bool, size)

    for i, bound := 3, int(math.Sqrt(float64(n))) + 1; i < bound; i += 2 {
        if sievedFactors[i >> 1] {
            continue
        }
        for j := i * i; j <= n; j += i << 1 {
            if !sievedFactors[j >> 1] {
                sievedFactors[j >> 1] = true
            }
        }
    }

    result := int64(1)

    for i := 2; i <= n; i <<= 1 {
        result <<= 1
    }

    var factor int
    for i := 1; i < size; i++ {
        if !sievedFactors[i] {
            factor = (i << 1) + 1
            for j := factor; j <= n; j *= factor {
                result *= int64(factor)
            }
        }
    }

    return result
}

func main() {
    var testCases int

    var n int

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
