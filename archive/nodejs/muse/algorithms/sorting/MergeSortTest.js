'use strict'

const MergeSort = require('muse/algorithms/sorting/MergeSort')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const Sequences = require('muse/util/Sequences')
const TestRunner = require('muse/util/TestRunner')

const testMergeSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  MergeSort.sort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

if (module === require.main) {
  TestRunner.pick(testMergeSort)
}
