package main

import (
    "fmt"
    "io"
    "math"
)

func draft(n int) int {
    if n == 1 {
        return 2
    }

    primes := make([]int, n)

    primes[0] = 2

    k := 1

    var isPrime bool
    var bound int

    for i := 3; i < 2147483647; i += 2 {
        isPrime = true
        bound = int(math.Sqrt(float64(i)))
        for j := 1; j < k; j++ {
            if primes[j] > bound {
                break
            }
            if i % primes[j] == 0 {
                isPrime = false
                break
            }
        }
        if isPrime {
            primes[k] = i
            k++
            if k == n {
                break
            }
        }
    }

    return primes[n - 1]
}

func process(n int) int {
    if n == 1 {
        return 2
    }

    var ceiling int

    if n < 6 {
        ceiling = (n << 1) + 1
    } else {
        cache := math.Log(float64(n))
        ceiling = int(float64(n) * (cache + math.Log(cache)))
    }

    size := (ceiling + 1) >> 1

    arr := make([]bool, size)

    for i, bound := 3, int(math.Sqrt(float64(ceiling))) + 1; i < bound; i += 2 {
        if arr[i >> 1] {
            continue
        }
        for j := i * i; j <= ceiling; j += i << 1 {
            if !arr[j >> 1] {
                arr[j >> 1] = true
            }
        }
    }

    k := 1
    for i := 1; i < size; i++ {
        if !arr[i] {
            k++
            if k == n {
                return (i << 1) + 1
            }
        }
    }

    return 0
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
