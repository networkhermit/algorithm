require "muse/algorithms/sorting/MergeSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testMergeSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    Sequences.isSorted(arr)
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testMergeSort())
end
