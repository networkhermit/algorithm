"use strict"

const LinkedStack = require("fun/vac/data_structures/LinkedStack")
const TestRunner  = require("fun/vac/util/TestRunner")

const testLinkedStack = () => {
    let size = 8192

    const stack = new LinkedStack.LinkedStack()

    for (let i = 1; i <= size; i++) {
        stack.push(i)
    }

    if (stack.size() !== size) {
        return false
    }

    for (let i = size; i > 0; i--) {
        if (stack.peek() !== i) {
            return false
        }
        stack.pop()
    }

    if (!stack.isEmpty()) {
        return false
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testLinkedStack())
}
