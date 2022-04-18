require "muse/algorithms/sorting/MergeSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testMergeSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  MergeSort.sort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

TestRunner.pick(testMergeSort) if __FILE__ == $PROGRAM_NAME
