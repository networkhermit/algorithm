from muse.algorithms.meta import Sorting
from muse.util import SequenceBuilder
from muse.util import Sequences
from muse.util import TestRunner

def testBubbleSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.bubbleSort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def testInsertionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.insertionSort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def testMergeSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.mergeSort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def testQuickSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.quickSort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def testSelectionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    Sorting.selectionSort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def main() -> None:
    TestRunner.parseTest(testBubbleSort())

    TestRunner.parseTest(testInsertionSort())

    TestRunner.parseTest(testMergeSort())

    TestRunner.parseTest(testQuickSort())

    TestRunner.parseTest(testSelectionSort())

if __name__ == "__main__":
    main()
