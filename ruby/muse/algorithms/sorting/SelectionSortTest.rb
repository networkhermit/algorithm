require "muse/algorithms/sorting/SelectionSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testSelectionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    SelectionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    Sequences.isSorted(arr)
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.pick(testSelectionSort())
end
