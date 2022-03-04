package main

import (
    "fmt"
    "io"
    "math"
)

func process(m int, n int) {
    if n == 1 {
        return
    }

    if m <= 2 {
        fmt.Println(2)
        m = 3
    }

    if n == 2 {
        return
    }

    m |= 1
    n = (n - 1) | 1

    cache := int(math.Sqrt(float64(n)))

    sievedFactors := make([]bool, (cache + 1) >> 1)

    for i, bound := 3, cache + 1; i < bound; i += 2 {
        if sievedFactors[i >> 1] {
            continue
        }
        for j := i * i; j < bound; j += i << 1 {
            if !sievedFactors[j >> 1] {
                sievedFactors[j >> 1] = true
            }
        }
    }

    nums := make([]bool, ((n - m) >> 1) + 1)

    var factor int
    var target int

    boost := m <= 9

    for i, length := 1, len(sievedFactors); i < length; i++ {
        if sievedFactors[i] {
            continue
        }

        factor = (i << 1) + 1

        if boost {
            target = factor * factor
        } else if m <= factor * factor {
            boost = true
            target = factor * factor
        } else {
            if m % factor == 0 {
                target = m
            } else {
                target = m + (factor - m % factor)
                if (target & 1) == 0 {
                    target += factor
                }
            }
        }

        for j := target; j <= n; j += factor << 1 {
            if !nums[(j - m) >> 1] {
                nums[(j - m) >> 1] = true
            }
        }
    }

    for i, v := range nums {
        if !v {
            fmt.Println(m + (i << 1))
        }
    }
}

func main() {
    var testCases int

    var m int
    var n int

    for {
        _, err := fmt.Scan(&testCases)
        if err == io.EOF {
            break
        }
        for i := 0; i < testCases; i++ {
            fmt.Scan(&m, &n)
            if i != 0 {
                fmt.Println()
            }
            process(m, n)
        }
    }
}
