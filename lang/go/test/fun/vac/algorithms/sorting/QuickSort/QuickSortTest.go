package main

import (
    "fun/vac/algorithms/sorting/QuickSort"
    "fun/vac/util/SequenceBuilder"
    "fun/vac/util/Sequences"
    "fun/vac/util/TestRunner"
)

func testQuickSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    QuickSort.Sort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func main() {
    TestRunner.ParseTest(testQuickSort())
}
