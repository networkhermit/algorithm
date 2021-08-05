require "muse/algorithms/sorting/BubbleSort"
require "muse/util/SequenceBuilder"
require "muse/util/Sequences"
require "muse/util/TestRunner"

def testBubbleSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    BubbleSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    Sequences.isSorted(arr)
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.pick(testBubbleSort())
end
