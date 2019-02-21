"use strict"

const LinkedQueue = require("fun/vac/data_structures/LinkedQueue")
const TestRunner  = require("fun/vac/util/TestRunner")

const testLinkedQueue = () => {
    let size = 8192

    const queue = new LinkedQueue.LinkedQueue()

    for (let i = 1; i <= size; i++) {
        queue.enqueue(i)
    }

    if (queue.size() !== size) {
        return false
    }

    for (let i = 1; i <= size; i++) {
        if (queue.peek() !== i) {
            return false
        }
        queue.dequeue()
    }

    if (!queue.isEmpty()) {
        return false
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testLinkedQueue())
}
