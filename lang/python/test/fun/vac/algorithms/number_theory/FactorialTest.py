from fun.vac.algorithms.number_theory import Factorial
from fun.vac.util import TestRunner

def testFactorial() -> bool:
    sample = [
        # fmt: off
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
        # fmt: on
    ]

    for i, _ in enumerate(sample):
        if Factorial.iterativeProcedure(sample[i][0]) != sample[i][1]:
            return False

    for i, _ in enumerate(sample):
        if Factorial.recursiveProcedure(sample[i][0]) != sample[i][1]:
            return False

    return True

def main() -> None:
    TestRunner.parseTest(testFactorial())

if __name__ == "__main__":
    main()
