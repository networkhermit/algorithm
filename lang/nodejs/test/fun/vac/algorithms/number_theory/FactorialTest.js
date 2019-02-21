"use strict"

const Factorial  = require("fun/vac/algorithms/number_theory/Factorial")
const TestRunner = require("fun/vac/util/TestRunner")

const testFactorial = () => {
    let mapping = [
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

    let instances = mapping.length

    for (let i = 0; i < instances; i++) {
        if (Factorial.iterativeProcedure(mapping[i][0]) !== mapping[i][1]) {
            return false
        }
    }

    for (let i = 0; i < instances; i++) {
        if (Factorial.recursiveProcedure(mapping[i][0]) !== mapping[i][1]) {
            return false
        }
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testFactorial())
}
