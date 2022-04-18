'use strict'

const Search = require('muse/algorithms/meta/Search')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const TestRunner = require('muse/util/TestRunner')

const testBinarySearch = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packIncreasing(arr)

  if (Search.binarySearch(arr, -1) !== size) {
    return false
  }

  if (Search.binarySearch(arr, 2147483647) !== size) {
    return false
  }

  for (let i = 0; i < size; i++) {
    if (Search.binarySearch(arr, arr[i]) !== i) {
      return false
    }
  }

  return true
}

const testLinearSearch = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packIncreasing(arr)

  if (Search.linearSearch(arr, -1) !== size) {
    return false
  }

  if (Search.linearSearch(arr, 2147483647) !== size) {
    return false
  }

  for (let i = 0; i < size; i++) {
    if (Search.linearSearch(arr, arr[i]) !== i) {
      return false
    }
  }

  return true
}

if (module === require.main) {
  TestRunner.pick(testBinarySearch)

  TestRunner.pick(testLinearSearch)
}
