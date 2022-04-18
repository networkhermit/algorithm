from muse.algorithms.sorting import MergeSort
from muse.util import SequenceBuilder, Sequences, TestRunner


def testMergeSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)


def main() -> None:
    TestRunner.pick(testMergeSort)


if __name__ == "__main__":
    main()
