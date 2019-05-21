"use strict"

let TestRunner_TestIndex = 0

exports.parseTest = (ok) => {
    if (ok) {
        console.log("✓ Test [%d] Passed", TestRunner_TestIndex)
    } else {
        console.error("✗ Test [%d] Failed", TestRunner_TestIndex)
    }

    TestRunner_TestIndex += 1
}
