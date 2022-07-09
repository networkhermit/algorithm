'use strict'

const ArrayQueue = require('muse/data_structures/ArrayQueue')
const TestRunner = require('muse/util/TestRunner')

const testArrayQueue = () => {
  const size = 8192

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

  return queue.capacity() === 0
}

if (module === require.main) {
  TestRunner.pick(testArrayQueue)
}
