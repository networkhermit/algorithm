require "fun/vac/algorithms/sorting/SelectionSort"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/Sequences"
require "fun/vac/util/TestRunner"

def testSelectionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    SelectionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    return Sequences.isSorted(arr)
end

if __FILE__ == $0
    TestRunner.parseTest(testSelectionSort())
end
