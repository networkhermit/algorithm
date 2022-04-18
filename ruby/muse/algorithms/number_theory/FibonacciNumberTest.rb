require "muse/algorithms/number_theory/FibonacciNumber"
require "muse/util/TestRunner"

def testFibonacciNumber
  sample = [
    [-31, 1_346_269],
    [-30, -832_040],
    [-29, 514_229],
    [-28, -317_811],
    [-27, 196_418],
    [-26, -121_393],
    [-25, 75_025],
    [-24, -46_368],
    [-23, 28_657],
    [-22, -17_711],
    [-21, 10_946],
    [-20, -6765],
    [-19, 4181],
    [-18, -2584],
    [-17, 1597],
    [-16, -987],
    [-15, 610],
    [-14, -377],
    [-13, 233],
    [-12, -144],
    [-11, 89],
    [-10, -55],
    [-9, 34],
    [-8, -21],
    [-7, 13],
    [-6, -8],
    [-5, 5],
    [-4, -3],
    [-3, 2],
    [-2, -1],
    [-1, 1],
    [0, 0],
    [1, 1],
    [2, 1],
    [3, 2],
    [4, 3],
    [5, 5],
    [6, 8],
    [7, 13],
    [8, 21],
    [9, 34],
    [10, 55],
    [11, 89],
    [12, 144],
    [13, 233],
    [14, 377],
    [15, 610],
    [16, 987],
    [17, 1597],
    [18, 2584],
    [19, 4181],
    [20, 6765],
    [21, 10_946],
    [22, 17_711],
    [23, 28_657],
    [24, 46_368],
    [25, 75_025],
    [26, 121_393],
    [27, 196_418],
    [28, 317_811],
    [29, 514_229],
    [30, 832_040],
    [31, 1_346_269]
  ]

  sample.each_index do |i|
    return false if FibonacciNumber.iterativeProcedure(sample[i][0]) != sample[i][1]
  end

  sample.each_index do |i|
    return false if FibonacciNumber.recursiveProcedure(sample[i][0]) != sample[i][1]
  end

  true
end

TestRunner.pick(testFibonacciNumber) if __FILE__ == $PROGRAM_NAME
