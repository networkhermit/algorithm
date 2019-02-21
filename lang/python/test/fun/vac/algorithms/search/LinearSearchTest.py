from fun.vac.algorithms.search import LinearSearch
from fun.vac.util import SequenceBuilder
from fun.vac.util import TestRunner

def testLinearSearch() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packIncreasing(arr)

    if LinearSearch.find(arr, -1) != size:
        return False

    if LinearSearch.find(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if LinearSearch.find(arr, v) != i:
            return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testLinearSearch())
