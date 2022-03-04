package main

import (
    "fmt"
    "io"
)

func process(n int, m int, a int) int64 {
    x := (n - 1) / a + 1
    y := (m - 1) / a + 1

    return int64(x) * int64(y)
}

func main() {
    var n int
    var m int
    var a int

    for {
        _, err := fmt.Scan(&n, &m, &a)
        if err == io.EOF {
            break
        }
        fmt.Println(process(n, m, a))
    }
}
