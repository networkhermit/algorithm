"use strict"

let TestRunnerItemIndex = 0

exports.parseTest = (ok) => {
    if (ok) {
        console.log("✓ Item [%d] PASSED", TestRunnerItemIndex)
    } else {
        console.error("✗ Item [%d] FAILED", TestRunnerItemIndex)
    }

    TestRunnerItemIndex++
}
