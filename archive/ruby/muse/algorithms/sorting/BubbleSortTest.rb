require "muse/algorithms/sorting/BubbleSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testBubbleSort
  size = 32_768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  BubbleSort.sort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

TestRunner.pick(testBubbleSort) if __FILE__ == $PROGRAM_NAME
