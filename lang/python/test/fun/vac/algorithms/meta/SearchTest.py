from fun.vac.algorithms.meta import Search
from fun.vac.util import SequenceBuilder
from fun.vac.util import TestRunner

def testBinarySearch() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packIncreasing(arr)

    if Search.binarySearch(arr, -1) != size:
        return False

    if Search.binarySearch(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if Search.binarySearch(arr, v) != i:
            return False

    return True

def testLinearSearch() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packIncreasing(arr)

    if Search.linearSearch(arr, -1) != size:
        return False

    if Search.linearSearch(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if Search.linearSearch(arr, v) != i:
            return False

    return True

if __name__ == "__main__":

    TestRunner.parseTest(testBinarySearch())

    TestRunner.parseTest(testLinearSearch())
