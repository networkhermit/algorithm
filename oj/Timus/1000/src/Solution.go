package main

import (
    "fmt"
    "io"
)

func main() {
    var a int
    var b int

    for {
        _, err := fmt.Scan(&a, &b)
        if err == io.EOF {
            break
        }
        fmt.Println(a + b)
    }
}
