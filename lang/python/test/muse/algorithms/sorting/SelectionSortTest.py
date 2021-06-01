from muse.algorithms.sorting import SelectionSort
from muse.util import SequenceBuilder
from muse.util import Sequences
from muse.util import TestRunner

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
    TestRunner.parseTest(testSelectionSort())

if __name__ == "__main__":
    main()
