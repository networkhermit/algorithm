"use strict"

const ArrayQueue = require("fun/vac/data_structures/ArrayQueue")
const TestRunner = require("fun/vac/util/TestRunner")

const testArrayQueue = () => {
    let size = 8192

    const queue = new ArrayQueue.ArrayQueue(0)

    for (let i = 1; i <= size; i++) {
        queue.enqueue(i)
    }

    queue.shrink()

    if (queue.size() !== size) {
        return false
    }

    if (queue.capacity() !== size) {
        return false
    }

    for (let i = 1; i <= size; i++) {
        if (queue.peek() !== i) {
            return false
        }
        queue.dequeue()
    }

    queue.shrink()

    if (!queue.isEmpty()) {
        return false
    }

    if (queue.capacity() !== 0) {
        return false
    }

    return true
}

if (module === require.main) {
    TestRunner.parseTest(testArrayQueue())
}
