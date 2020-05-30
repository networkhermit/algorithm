require "fun/vac/algorithms/number_theory/Factorial"
require "fun/vac/util/TestRunner"

def testFactorial()
    sample = [
        [ 0,           1],
        [ 1,           1],
        [ 2,           2],
        [ 3,           6],
        [ 4,          24],
        [ 5,         120],
        [ 6,         720],
        [ 7,        5040],
        [ 8,       40320],
        [ 9,      362880],
        [10,   3_628_800],
        [11,  39_916_800],
        [12, 479_001_600],
    ]

    sample.each_index do |i|
        if Factorial.iterativeProcedure(sample[i][0]) != sample[i][1]
            return false
        end
    end

    sample.each_index do |i|
        if Factorial.recursiveProcedure(sample[i][0]) != sample[i][1]
            return false
        end
    end

    true
end

if __FILE__ == $PROGRAM_NAME
    TestRunner.parseTest(testFactorial())
end
