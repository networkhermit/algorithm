package main

import (
    "fmt"
    "io"
)

func process(sentinel int64) int64 {
    sum := int64(0)

    prev := int64(0)
    curr := int64(2)

    for curr < sentinel {
        sum += curr
        prev, curr = curr, prev + (curr << 2)
    }

    return sum
}

func main() {
    var testCases int

    var sentinel int64

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
