"use strict"

const MergeSort       = require("vac.fun/algorithms/sorting/MergeSort")
const SequenceBuilder = require("vac.fun/util/SequenceBuilder")
const Sequences       = require("vac.fun/util/Sequences")
const TestRunner      = require("vac.fun/util/TestRunner")

const testMergeSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

if (module === require.main) {
    TestRunner.parseTest(testMergeSort())
}