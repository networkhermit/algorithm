package main

import (
    "vac.fun/algorithms/meta/Search"
    "vac.fun/util/SequenceBuilder"
    "vac.fun/util/TestRunner"
)

func testBinarySearch() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if Search.BinarySearch(arr, -1) != size {
        return false
    }

    if Search.BinarySearch(arr, 2147483647) != size {
        return false
    }

    for i, v := range arr {
        if Search.BinarySearch(arr, v) != i {
            return false
        }
    }

    return true
}

func testLinearSearch() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if Search.LinearSearch(arr, -1) != size {
        return false
    }

    if Search.LinearSearch(arr, 2147483647) != size {
        return false
    }

    for i, v := range arr {
        if Search.LinearSearch(arr, v) != i {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testBinarySearch())

    TestRunner.ParseTest(testLinearSearch())
}
