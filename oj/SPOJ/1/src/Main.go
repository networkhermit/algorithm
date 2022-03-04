package main

import (
    "fmt"
    "io"
)

func main() {
    var n int

    for {
        _, err := fmt.Scan(&n)
        if err == io.EOF {
            break
        }
        if n == 42 {
            break
        }
        fmt.Println(n)
    }
}
