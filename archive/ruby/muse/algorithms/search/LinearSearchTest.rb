require "muse/algorithms/search/LinearSearch"
require "muse/util/SequenceBuilder"
require "muse/util/TestRunner"

def testLinearSearch
  size = 32768

  arr = Array.new(size)
  SequenceBuilder.packIncreasing(arr)

  return false if LinearSearch.find(arr, -1) != size

  return false if LinearSearch.find(arr, 2_147_483_647) != size

  arr.each_index do |i|
    return false if LinearSearch.find(arr, arr[i]) != i
  end

  true
end

TestRunner.pick(testLinearSearch) if __FILE__ == $PROGRAM_NAME
