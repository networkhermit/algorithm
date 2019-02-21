package main

import (
    "fun/vac/algorithms/meta/Sorting"
    "fun/vac/util/SequenceBuilder"
    "fun/vac/util/Sequences"
    "fun/vac/util/TestRunner"
)

func testBubbleSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sorting.BubbleSort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func testInsertionSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sorting.InsertionSort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func testMergeSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sorting.MergeSort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func testQuickSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sorting.QuickSort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func testSelectionSort() bool {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sorting.SelectionSort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        return false
    }

    if !Sequences.IsSorted(arr) {
        return false
    }

    return true
}

func main() {

    TestRunner.ParseTest(testBubbleSort())

    TestRunner.ParseTest(testInsertionSort())

    TestRunner.ParseTest(testMergeSort())

    TestRunner.ParseTest(testQuickSort())

    TestRunner.ParseTest(testSelectionSort())
}
