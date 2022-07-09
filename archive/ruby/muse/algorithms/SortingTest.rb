require "muse/algorithms/Sorting"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testBubbleSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  Sorting.bubbleSort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

def testInsertionSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  Sorting.insertionSort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

def testMergeSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  Sorting.mergeSort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

def testQuickSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  Sorting.quickSort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

def testSelectionSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  Sorting.selectionSort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

if __FILE__ == $PROGRAM_NAME
  TestRunner.pick(testBubbleSort)

  TestRunner.pick(testInsertionSort)

  TestRunner.pick(testMergeSort)

  TestRunner.pick(testQuickSort)

  TestRunner.pick(testSelectionSort)
end
