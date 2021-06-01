package main

import (
    "muse/algorithms/sorting/QuickSort"
    "muse/util/SequenceBuilder"
    "muse/util/Sequences"
    "muse/util/TestRunner"
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

    return Sequences.IsSorted(arr)
}

func main() {
    TestRunner.ParseTest(testQuickSort())
}
