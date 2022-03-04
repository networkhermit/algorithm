package main

import (
    "fmt"
    "io"
)

func draft(n int) bool {
    k := 0
    arr := make([]int, 10)

    for n != 0 {
        arr[k] = n % 10
        n /= 10
        k++
    }

    for i, j, bound := 0, k - 1, (k >> 1) + 1; i < bound; i, j = i + 1, j - 1 {
        if arr[i] != arr[j] {
            return false
        }
    }

    return true
}

func isPalindrome(n int) bool {
    original := n

    reversed := 0
    for n != 0 {
        reversed = reversed * 10 + n % 10
        n /= 10
    }

    return original == reversed
}

func exhaustiveMatch(n int) bool {
    var product int

    for i := 999; i > 99; i-- {
        for j := i; j > 99; j-- {
            product = i * j
            if product == n {
                return true
            } else if product < n {
                break
            }
        }
    }

    return false
}

func process(sentinel int) int {
    for i := sentinel - 1; i > 10000; i-- {
        if isPalindrome(i) && exhaustiveMatch(i) {
            return i
        }
    }
    return 0
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
