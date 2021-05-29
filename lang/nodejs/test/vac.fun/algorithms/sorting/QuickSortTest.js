"use strict"

const QuickSort       = require("vac.fun/algorithms/sorting/QuickSort")
const SequenceBuilder = require("vac.fun/util/SequenceBuilder")
const Sequences       = require("vac.fun/util/Sequences")
const TestRunner      = require("vac.fun/util/TestRunner")

const testQuickSort = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packRandom(arr)

    let checksum = Sequences.parityChecksum(arr)

    QuickSort.sort(arr)

    if (Sequences.parityChecksum(arr) !== checksum) {
        return false
    }

    return Sequences.isSorted(arr)
}

if (module === require.main) {
    TestRunner.parseTest(testQuickSort())
}
