package main

import (
    "muse/algorithms/sorting/InsertionSort"
    "muse/util/SequenceBuilder"
    "muse/util/Sequences"
    "muse/util/TestRunner"
)

func testInsertionSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    InsertionSort.Sort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    return Sequences.IsSorted(arr)
}

func main() {
    TestRunner.ParseTest(testInsertionSort())
}
