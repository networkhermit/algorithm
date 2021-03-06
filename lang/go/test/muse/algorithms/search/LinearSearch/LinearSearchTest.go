package main

import (
    "muse/algorithms/search/LinearSearch"
    "muse/util/SequenceBuilder"
    "muse/util/TestRunner"
)

func testLinearSearch() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if LinearSearch.Find(arr, -1) != size {
        return false
    }

    if LinearSearch.Find(arr, 2147483647) != size {
        return false
    }

    for i, v := range arr {
        if LinearSearch.Find(arr, v) != i {
            return false
        }
    }

    return true
}

func main() {
    TestRunner.ParseTest(testLinearSearch())
}
