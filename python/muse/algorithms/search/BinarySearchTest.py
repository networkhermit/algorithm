from muse.algorithms.search import BinarySearch
from muse.util import SequenceBuilder
from muse.util import TestRunner

def testBinarySearch() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packIncreasing(arr)

    if BinarySearch.find(arr, -1) != size:
        return False

    if BinarySearch.find(arr, 2_147_483_647) != size:
        return False

    for i, v in enumerate(arr):
        if BinarySearch.find(arr, v) != i:
            return False

    return True

def main() -> None:
    TestRunner.parseTest(testBinarySearch())

if __name__ == "__main__":
    main()
