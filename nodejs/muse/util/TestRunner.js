"use strict"

let TestRunnerItemIndex = 0

exports.pick = (func) => {
    if (func()) {
        console.log("✓ Item [%d] PASSED", TestRunnerItemIndex)
    } else {
        console.error("✗ Item [%d] FAILED", TestRunnerItemIndex)
    }

    TestRunnerItemIndex++
}
