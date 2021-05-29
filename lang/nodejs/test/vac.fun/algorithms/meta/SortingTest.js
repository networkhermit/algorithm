"use strict"

const Sorting         = require("vac.fun/algorithms/meta/Sorting")
const SequenceBuilder = require("vac.fun/util/SequenceBuilder")
const Sequences       = require("vac.fun/util/Sequences")
const TestRunner      = require("vac.fun/util/TestRunner")

const testBubbleSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.bubbleSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

const testInsertionSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.insertionSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

const testMergeSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.mergeSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

const testQuickSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.quickSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

const testSelectionSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.selectionSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

if (module === require.main) {
    TestRunner.parseTest(testBubbleSort())

    TestRunner.parseTest(testInsertionSort())

    TestRunner.parseTest(testMergeSort())

    TestRunner.parseTest(testQuickSort())

    TestRunner.parseTest(testSelectionSort())
}
