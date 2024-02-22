require "muse/algorithms/search/BinarySearch"
require "muse/util/SequenceBuilder"
require "muse/util/TestRunner"

def testBinarySearch
  size = 32768

  arr = Array.new(size)
  SequenceBuilder.packIncreasing(arr)

  return false if BinarySearch.find(arr, -1) != size

  return false if BinarySearch.find(arr, 2_147_483_647) != size

  arr.each_index do |i|
    return false if BinarySearch.find(arr, arr[i]) != i
  end

  true
end

TestRunner.pick(testBinarySearch) if __FILE__ == $PROGRAM_NAME
