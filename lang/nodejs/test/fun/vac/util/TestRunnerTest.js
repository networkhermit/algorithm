"use strict"

const TestRunner = require("fun/vac/util/TestRunner")

const testParseTest = () => {
    for (let i = 0; i < 10; i++) {
        TestRunner.parseTest((i & 1) !== 0)
    }
}

if (module === require.main) {
    testParseTest()
}
