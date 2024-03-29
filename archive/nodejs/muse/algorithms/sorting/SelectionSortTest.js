'use strict'

const SelectionSort = require('muse/algorithms/sorting/SelectionSort')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const Sequences = require('muse/util/Sequences')
const TestRunner = require('muse/util/TestRunner')

const testSelectionSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  SelectionSort.sort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

if (module === require.main) {
  TestRunner.pick(testSelectionSort)
}
