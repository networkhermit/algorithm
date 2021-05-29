require "vac.fun/algorithms/sorting/InsertionSort"
require "vac.fun/util/SequenceBuilder"
require "vac.fun/util/Sequences"
require "vac.fun/util/TestRunner"

def testInsertionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    InsertionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    Sequences.isSorted(arr)
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testInsertionSort())
end
