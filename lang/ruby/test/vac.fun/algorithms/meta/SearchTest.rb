require "vac.fun/algorithms/meta/Search"
require "vac.fun/util/SequenceBuilder"
require "vac.fun/util/TestRunner"

def testBinarySearch()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packIncreasing(arr)

    if Search.binarySearch(arr, -1) != size
        return false
    end

    if Search.binarySearch(arr, 2_147_483_647) != size
        return false
    end

    arr.each_index do |i|
        if Search.binarySearch(arr, arr[i]) != i
            return false
        end
    end

    true
end

def testLinearSearch()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packIncreasing(arr)

    if Search.linearSearch(arr, -1) != size
        return false
    end

    if Search.linearSearch(arr, 2_147_483_647) != size
        return false
    end

    arr.each_index do |i|
        if Search.linearSearch(arr, arr[i]) != i
            return false
        end
    end

    true
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testBinarySearch())

    TestRunner.parseTest(testLinearSearch())
end
