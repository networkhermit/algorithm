'use strict'

const BinarySearch = require('muse/algorithms/search/BinarySearch')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const TestRunner = require('muse/util/TestRunner')

const testBinarySearch = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packIncreasing(arr)

  if (BinarySearch.find(arr, -1) !== size) {
    return false
  }

  if (BinarySearch.find(arr, 2147483647) !== size) {
    return false
  }

  for (let i = 0; i < size; i++) {
    if (BinarySearch.find(arr, arr[i]) !== i) {
      return false
    }
  }

  return true
}

if (module === require.main) {
  TestRunner.pick(testBinarySearch)
}
