'use strict'

const CircularlyLinkedList = require('muse/data_structures/CircularlyLinkedList')
const TestRunner = require('muse/util/TestRunner')

const testCircularlyLinkedList = () => {
  const size = 8192

  const list = new CircularlyLinkedList.CircularlyLinkedList()

  for (let i = 1; i <= size; i++) {
    list.append(i)
  }

  if (list.size() !== size) {
    return false
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== i + 1) {
      return false
    }
  }

  for (let i = 0; i < size; i++) {
    list.set(i, size - i)
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== size - i) {
      return false
    }
  }

  for (let i = 0, j = size - 1; i < j; i++, j--) {
    const x = list.get(i)
    const y = list.get(j)

    list.remove(i)
    list.insert(i, y)
    list.remove(j)
    list.insert(j, x)
  }

  for (let i = 0; i < size; i++) {
    if (list.get(i) !== i + 1) {
      return false
    }
  }

  for (let i = size; i >= 1; i--) {
    if (list.back() !== i) {
      return false
    }
    list.eject()
  }

  return list.isEmpty()
}

if (module === require.main) {
  TestRunner.pick(testCircularlyLinkedList)
}
