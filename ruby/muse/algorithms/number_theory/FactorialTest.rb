require "muse/algorithms/number_theory/Factorial"
require "muse/util/TestRunner"

def testFactorial
  sample = [
    [0, 1],
    [1, 1],
    [2, 2],
    [3, 6],
    [4, 24],
    [5, 120],
    [6, 720],
    [7, 5040],
    [8, 40_320],
    [9, 362_880],
    [10, 3_628_800],
    [11, 39_916_800],
    [12, 479_001_600]
  ]

  sample.each_index do |i|
    return false if Factorial.iterativeProcedure(sample[i][0]) != sample[i][1]
  end

  sample.each_index do |i|
    return false if Factorial.recursiveProcedure(sample[i][0]) != sample[i][1]
  end

  true
end

TestRunner.pick(testFactorial) if __FILE__ == $PROGRAM_NAME
