package main

import (
    "fmt"
    "io"
)

func draft(n int) int64 {
    v := int64(n)
    result := int64(0)

    for i := int64(1); i < v; i++ {
        result += i * (v - i) * (i + 1 + v)
    }

    return result
}

func process(n int) int64 {
    v := int64(n)
    return (v * (v + 1) >> 1) * (v - 1) / 3 * (3 * v + 2) >> 1
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
