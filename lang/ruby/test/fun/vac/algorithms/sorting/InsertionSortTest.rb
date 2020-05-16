require "fun/vac/algorithms/sorting/InsertionSort"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/Sequences"
require "fun/vac/util/TestRunner"

def testInsertionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    InsertionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    return Sequences.isSorted(arr)
end

if __FILE__ == $0
    TestRunner.parseTest(testInsertionSort())
end
