"use strict"

const BinarySearch    = require("vac.fun/algorithms/search/BinarySearch")
const SequenceBuilder = require("vac.fun/util/SequenceBuilder")
const TestRunner      = require("vac.fun/util/TestRunner")

const testBinarySearch = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packIncreasing(arr)

    if (BinarySearch.find(arr, -1) !== size) {
        return false
    }

    if (BinarySearch.find(arr, 2147483647) !== size) {
        return false
    }

    for (let i = 0; i < size; i++) {
        if (BinarySearch.find(arr, arr[i]) !== i) {
            return false
        }
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testBinarySearch())
}
