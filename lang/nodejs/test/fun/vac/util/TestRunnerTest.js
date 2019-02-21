"use strict"

const TestRunner = require("fun/vac/util/TestRunner")

const testParseTest = () => {
    for (let i = 0; i < 10; i++) {
        if ((i & 1) === 0) {
            TestRunner.parseTest(false)
        } else {
            TestRunner.parseTest(true)
        }
    }
}

if (module === require.main) {
    testParseTest()
}
