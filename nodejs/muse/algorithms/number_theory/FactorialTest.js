"use strict"

const Factorial  = require("muse/algorithms/number_theory/Factorial")
const TestRunner = require("muse/util/TestRunner")

const testFactorial = () => {
    let sample = [
        [ 0,         1],
        [ 1,         1],
        [ 2,         2],
        [ 3,         6],
        [ 4,        24],
        [ 5,       120],
        [ 6,       720],
        [ 7,      5040],
        [ 8,     40320],
        [ 9,    362880],
        [10,   3628800],
        [11,  39916800],
        [12, 479001600],
    ]

    for (let i = 0, size = sample.length; i < size; i++) {
        if (Factorial.iterativeProcedure(sample[i][0]) !== sample[i][1]) {
            return false
        }
    }

    for (let i = 0, size = sample.length; i < size; i++) {
        if (Factorial.recursiveProcedure(sample[i][0]) !== sample[i][1]) {
            return false
        }
    }

    return true
}

if (module === require.main) {
    TestRunner.pick(testFactorial)
}
