'use strict'

const QuickSort = require('muse/algorithms/sorting/QuickSort')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const Sequences = require('muse/util/Sequences')
const TestRunner = require('muse/util/TestRunner')

const testQuickSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  QuickSort.sort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

if (module === require.main) {
  TestRunner.pick(testQuickSort)
}
