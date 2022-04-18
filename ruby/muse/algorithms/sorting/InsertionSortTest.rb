require "muse/algorithms/sorting/InsertionSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testInsertionSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  InsertionSort.sort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

TestRunner.pick(testInsertionSort) if __FILE__ == $PROGRAM_NAME
