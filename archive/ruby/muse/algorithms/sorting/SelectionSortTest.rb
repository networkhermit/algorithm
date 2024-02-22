require "muse/algorithms/sorting/SelectionSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testSelectionSort
  size = 32768

  arr = Array.new(size)
  SequenceBuilder.packRandom(arr)

  checksum = Sequences.parityChecksum(arr)

  SelectionSort.sort(arr)

  return false if Sequences.parityChecksum(arr) != checksum

  Sequences.isSorted(arr)
end

TestRunner.pick(testSelectionSort) if __FILE__ == $PROGRAM_NAME
