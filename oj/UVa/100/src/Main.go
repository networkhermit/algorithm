package main

import (
    "fmt"
    "io"
)

func process(left int, right int) int {
    var n int

    var cycleLength int

    maxCycleLength := 0

    for i := left; i <= right; i++ {
        cycleLength = 0

        n = i

        for {
            cycleLength++
            if n == 1 {
                if cycleLength > maxCycleLength {
                    maxCycleLength = cycleLength
                }
                break
            }
            if (n & 1) == 0 {
                n >>= 1
            } else {
                n = (n << 1) + n + 1
            }
        }
    }

    return maxCycleLength
}

func main() {
    var m int
    var n int

    for {
        _, err := fmt.Scan(&m, &n)
        if err == io.EOF {
            break
        }
        if m < n {
            fmt.Printf("%d %d %d\n", m, n, process(m, n))
        } else {
            fmt.Printf("%d %d %d\n", m, n, process(n, m))
        }
    }
}
