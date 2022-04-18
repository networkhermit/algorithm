from muse.algorithms.meta import Search
from muse.util import SequenceBuilder, TestRunner


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


def main() -> None:
    TestRunner.pick(testBinarySearch)

    TestRunner.pick(testLinearSearch)


if __name__ == "__main__":
    main()
