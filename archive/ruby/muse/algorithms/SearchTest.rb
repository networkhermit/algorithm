require "muse/algorithms/Search"
require "muse/util/SequenceBuilder"
require "muse/util/TestRunner"

def testBinarySearch
  size = 32768

  arr = Array.new(size)
  SequenceBuilder.packIncreasing(arr)

  return false if Search.binarySearch(arr, -1) != size

  return false if Search.binarySearch(arr, 2_147_483_647) != size

  arr.each_index do |i|
    return false if Search.binarySearch(arr, arr[i]) != i
  end

  true
end

def testLinearSearch
  size = 32768

  arr = Array.new(size)
  SequenceBuilder.packIncreasing(arr)

  return false if Search.linearSearch(arr, -1) != size

  return false if Search.linearSearch(arr, 2_147_483_647) != size

  arr.each_index do |i|
    return false if Search.linearSearch(arr, arr[i]) != i
  end

  true
end

if __FILE__ == $PROGRAM_NAME
  TestRunner.pick(testBinarySearch)

  TestRunner.pick(testLinearSearch)
end
