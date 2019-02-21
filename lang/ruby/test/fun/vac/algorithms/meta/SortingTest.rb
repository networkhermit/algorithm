require "fun/vac/algorithms/meta/Sorting"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/Sequences"
require "fun/vac/util/TestRunner"

def testBubbleSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.bubbleSort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

def testInsertionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.insertionSort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

def testMergeSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.mergeSort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

def testQuickSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.quickSort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

def testSelectionSort()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.selectionSort(arr)

    if Sequences.parityChecksum(arr) != checksum
        return false
    end

    unless Sequences.isSorted(arr)
        return false
    end

    return true
end

if __FILE__ == $0

    TestRunner.parseTest(testBubbleSort())

    TestRunner.parseTest(testInsertionSort())

    TestRunner.parseTest(testMergeSort())

    TestRunner.parseTest(testQuickSort())

    TestRunner.parseTest(testSelectionSort())
end
