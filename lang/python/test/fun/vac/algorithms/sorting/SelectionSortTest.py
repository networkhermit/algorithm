from fun.vac.algorithms.sorting import SelectionSort
from fun.vac.util import SequenceBuilder
from fun.vac.util import Sequences
from fun.vac.util import TestRunner

def testSelectionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    SelectionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    if not Sequences.isSorted(arr):
        return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testSelectionSort())
