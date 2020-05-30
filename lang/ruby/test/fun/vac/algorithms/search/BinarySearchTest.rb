require "fun/vac/algorithms/search/BinarySearch"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/TestRunner"

def testBinarySearch()
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packIncreasing(arr)

    if BinarySearch.find(arr, -1) != size
        return false
    end

    if BinarySearch.find(arr, 2_147_483_647) != size
        return false
    end

    arr.each_index do |i|
        if BinarySearch.find(arr, arr[i]) != i
            return false
        end
    end

    true
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testBinarySearch())
end
