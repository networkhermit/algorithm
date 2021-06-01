package main

import (
    "muse/algorithms/search/BinarySearch"
    "muse/util/SequenceBuilder"
    "muse/util/TestRunner"
)

func testBinarySearch() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if BinarySearch.Find(arr, -1) != size {
        return false
    }

    if BinarySearch.Find(arr, 2147483647) != size {
        return false
    }

    for i, v := range arr {
        if BinarySearch.Find(arr, v) != i {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testBinarySearch())
}
