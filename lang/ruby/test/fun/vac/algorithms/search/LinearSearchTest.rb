require "fun/vac/algorithms/search/LinearSearch"
require "fun/vac/util/SequenceBuilder"
require "fun/vac/util/TestRunner"

def testLinearSearch
    size = 32768

    arr = Array.new(size)
    SequenceBuilder.packIncreasing(arr)

    if LinearSearch.find(arr, -1) != size
        return false
    end

    if LinearSearch.find(arr, 2_147_483_647) != size
        return false
    end

    arr.each_index do |i|
        if LinearSearch.find(arr, arr[i]) != i
            return false
        end
    end

    return true
end

if __FILE__ == $0
    TestRunner.parseTest(testLinearSearch())
end
