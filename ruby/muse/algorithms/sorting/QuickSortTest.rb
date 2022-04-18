require "muse/algorithms/sorting/QuickSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testQuickSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  QuickSort.sort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

TestRunner.pick(testQuickSort) if __FILE__ == $PROGRAM_NAME
