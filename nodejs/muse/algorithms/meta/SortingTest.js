'use strict'

const Sorting = require('muse/algorithms/meta/Sorting')
const SequenceBuilder = require('muse/util/SequenceBuilder')
const Sequences = require('muse/util/Sequences')
const TestRunner = require('muse/util/TestRunner')

const testBubbleSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  Sorting.bubbleSort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

const testInsertionSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  Sorting.insertionSort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

const testMergeSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  Sorting.mergeSort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

const testQuickSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  Sorting.quickSort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

const testSelectionSort = () => {
  const size = 32768

  const arr = new Array(size)
  SequenceBuilder.packRandom(arr)

  const checksum = Sequences.parityChecksum(arr)

  Sorting.selectionSort(arr)

  if (Sequences.parityChecksum(arr) !== checksum) {
    return false
  }

  return Sequences.isSorted(arr)
}

if (module === require.main) {
  TestRunner.pick(testBubbleSort)

  TestRunner.pick(testInsertionSort)

  TestRunner.pick(testMergeSort)

  TestRunner.pick(testQuickSort)

  TestRunner.pick(testSelectionSort)
}
