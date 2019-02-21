from fun.vac.algorithms.sorting import QuickSort
from fun.vac.util import SequenceBuilder
from fun.vac.util import Sequences
from fun.vac.util import TestRunner

def testQuickSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    QuickSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    if not Sequences.isSorted(arr):
        return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testQuickSort())
