require "fun/vac/algorithms/sorting/MergeSort"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/Sequences"
require "fun/vac/util/TestRunner"

def testMergeSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

if __FILE__ == $0
    TestRunner.parseTest(testMergeSort())
end
