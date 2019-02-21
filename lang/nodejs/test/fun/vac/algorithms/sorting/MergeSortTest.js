"use strict"

const MergeSort       = require("fun/vac/algorithms/sorting/MergeSort")
const SequenceBuilder = require("fun/vac/util/SequenceBuilder")
const Sequences       = require("fun/vac/util/Sequences")
const TestRunner      = require("fun/vac/util/TestRunner")

const testMergeSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    if (!Sequences.isSorted(arr)) {
        return false
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testMergeSort())
}
