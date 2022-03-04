package main

import (
    "fmt"
    "io"
)

func bruteForce(sentinel int) int64 {
    sum := int64(0)
    for i := 0; i < sentinel; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += int64(i)
        }
    }
    return sum
}

func arithmeticSeries(sentinel int, unit int) int64 {
    n := int64((sentinel - 1) / unit)

    return int64(unit) * (n * (1 + n) >> 1)
}

func process(sentinel int) int64 {
    sum := int64(0)

    sum += arithmeticSeries(sentinel, 3)
    sum -= arithmeticSeries(sentinel, 15)
    sum += arithmeticSeries(sentinel, 5)

    return sum
}

func main() {
    var testCases int

    var sentinel int

    for {
        _, err := fmt.Scan(&testCases)
        if err == io.EOF {
            break
        }
        for i := 0; i < testCases; i++ {
            fmt.Scan(&sentinel)
            fmt.Println(process(sentinel))
        }
    }
}
