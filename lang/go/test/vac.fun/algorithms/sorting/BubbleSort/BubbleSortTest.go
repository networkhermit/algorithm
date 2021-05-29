package main

import (
    "vac.fun/algorithms/sorting/BubbleSort"
    "vac.fun/util/SequenceBuilder"
    "vac.fun/util/Sequences"
    "vac.fun/util/TestRunner"
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
