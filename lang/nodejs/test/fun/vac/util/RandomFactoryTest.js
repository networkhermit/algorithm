"use strict"

const RandomFactory = require("fun/vac/util/RandomFactory")
const TestRunner    = require("fun/vac/util/TestRunner")

const testGenerateEven = () => {
    RandomFactory.launch()

    let value = 0
    for (let i = 0; i < 8192; i++) {
        value = RandomFactory.generateEven()
        if ((value & 1) !== 0) {
            return false
        }
    }

    return true
}

const testGenerateOdd = () => {
    RandomFactory.launch()

    let value = 0
    for (let i = 0; i < 8192; i++) {
        value = RandomFactory.generateOdd()
        if ((value & 1) === 0) {
            return false
        }
    }

    return true
}

if (module === require.main) {

    TestRunner.parseTest(testGenerateEven())

    TestRunner.parseTest(testGenerateOdd())
}
