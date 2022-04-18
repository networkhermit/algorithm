from muse.algorithms.sorting import QuickSort
from muse.util import SequenceBuilder, Sequences, TestRunner


def testQuickSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    QuickSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)


def main() -> None:
    TestRunner.pick(testQuickSort)


if __name__ == "__main__":
    main()
