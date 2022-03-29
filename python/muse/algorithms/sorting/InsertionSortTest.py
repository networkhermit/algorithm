from muse.algorithms.sorting import InsertionSort
from muse.util import SequenceBuilder
from muse.util import Sequences
from muse.util import TestRunner


def testInsertionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    InsertionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)


def main() -> None:
    TestRunner.pick(testInsertionSort)


if __name__ == "__main__":
    main()
