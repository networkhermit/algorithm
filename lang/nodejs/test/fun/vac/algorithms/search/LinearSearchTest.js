"use strict"

const LinearSearch    = require("fun/vac/algorithms/search/LinearSearch")
const SequenceBuilder = require("fun/vac/util/SequenceBuilder")
const TestRunner      = require("fun/vac/util/TestRunner")

const testLinearSearch = () => {
    let size = 32768

    const arr = new Array(size)
    SequenceBuilder.packIncreasing(arr)

    if (LinearSearch.find(arr, -1) !== size) {
        return false
    }

    if (LinearSearch.find(arr, 2147483647) !== size) {
        return false
    }

    for (let i = 0; i < size; i++) {
        if (LinearSearch.find(arr, arr[i]) !== i) {
            return false
        }
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testLinearSearch())
}
