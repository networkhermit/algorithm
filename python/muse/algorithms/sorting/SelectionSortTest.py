from muse.algorithms.sorting import SelectionSort
from muse.util import SequenceBuilder, Sequences, TestRunner


def testSelectionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    SelectionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)


def main() -> None:
    TestRunner.pick(testSelectionSort)


if __name__ == "__main__":
    main()
