"use strict"

const Sorting         = require("fun/vac/algorithms/meta/Sorting")
const SequenceBuilder = require("fun/vac/util/SequenceBuilder")
const Sequences       = require("fun/vac/util/Sequences")
const TestRunner      = require("fun/vac/util/TestRunner")

const testBubbleSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    Sorting.bubbleSort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
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

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
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

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
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

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
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

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
}

if (module === require.main) {

    TestRunner.parseTest(testBubbleSort())

    TestRunner.parseTest(testInsertionSort())

    TestRunner.parseTest(testMergeSort())

    TestRunner.parseTest(testQuickSort())

    TestRunner.parseTest(testSelectionSort())
}
