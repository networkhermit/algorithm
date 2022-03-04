package main

import (
    "fmt"
    "io"
    "math"
)

func main() {
    s := make([]int64, 0)

    var n int64

    for {
        _, err := fmt.Scan(&n)
        if err == io.EOF {
            break
        }
        s = append(s, n)
    }

    for i := len(s); i > 0; i-- {
        fmt.Printf("%.4f\n", math.Sqrt(float64(s[i - 1])))
    }
}
