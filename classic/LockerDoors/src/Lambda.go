package main

import (
    "fmt"
    "math"
)

func bruteForce(n int) []bool {
    arr := make([]bool, n)

    for i := 1; i <= n; i++ {
        for j := i - 1; j < n; j += i {
            arr[j] = !arr[j]
        }
    }

    return arr
}

func process(n int) []bool {
    arr := make([]bool, n)

    var temp int
    for i := 1; i <= n; i++ {
        temp = int(math.Sqrt(float64(i)))
        arr[i - 1] = temp * temp == i
    }

    return arr
}

func main() {
    const N = 100

    arr := process(N)

    for i, v := range arr {
        if v {
            fmt.Printf("%d\t  OPEN\n", i + 1)
        }
    }
}
