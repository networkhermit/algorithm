"use strict"

const LinkedStack = require("muse/data_structures/LinkedStack")
const TestRunner  = require("muse/util/TestRunner")

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

    return stack.isEmpty()
}

if (module === require.main) {
    TestRunner.pick(testLinkedStack)
}
