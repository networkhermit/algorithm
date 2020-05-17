package Sequences

import (
    "fmt"

    "fun/vac/algorithms/sorting/QuickSort"
    "fun/vac/util/RandomFactory"
)

func Inspect(arr []int) {
    fmt.Println("[")
    for i, v := range arr {
        fmt.Printf("\t#%04X  ->  %d\n", i, v)
    }
    fmt.Println("]")
}

func IsSorted(arr []int) bool {
    for i, length := 1, len(arr); i < length; i++ {
        if arr[i] < arr[i - 1] {
            return false
        }
    }

    return true
}

func ParityChecksum(arr []int) int {
    checksum := 0

    for _, v := range arr {
        checksum ^= v
    }

    return checksum
}

func Reverse(arr []int) {
    var k int

    for i, bound, length := 0, len(arr) >> 1, len(arr); i < bound; i++ {
        k = length - i - 1
        arr[i], arr[k] = arr[k], arr[i]
    }
}

func Shuffle(arr []int) {
    var k int

    RandomFactory.Seed()
    for i, length := 0, len(arr); i < length; i++ {
        k = RandomFactory.IntegerN(i, length)
        arr[i], arr[k] = arr[k], arr[i]
    }
}

func Sort(arr []int) {
    QuickSort.Sort(arr)
}
