from muse.algorithms.search import LinearSearch
from muse.util import SequenceBuilder
from muse.util import TestRunner


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


def main() -> None:
    TestRunner.pick(testLinearSearch)


if __name__ == "__main__":
    main()
