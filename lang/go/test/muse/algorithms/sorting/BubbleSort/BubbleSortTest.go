package main

import (
    "muse/algorithms/sorting/BubbleSort"
    "muse/util/SequenceBuilder"
    "muse/util/Sequences"
    "muse/util/TestRunner"
)

func testBubbleSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    BubbleSort.Sort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    return Sequences.IsSorted(arr)
}

func main() {
    TestRunner.ParseTest(testBubbleSort())
}
