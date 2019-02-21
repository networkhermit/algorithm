from fun.vac.algorithms.number_theory import Factorial
from fun.vac.util import TestRunner

def testFactorial() -> bool:
    mapping = [
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

    instances = len(mapping)

    for i in range(instances):
        if Factorial.iterativeProcedure(mapping[i][0]) != mapping[i][1]:
            return False

    for i in range(instances):
        if Factorial.recursiveProcedure(mapping[i][0]) != mapping[i][1]:
            return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testFactorial())
