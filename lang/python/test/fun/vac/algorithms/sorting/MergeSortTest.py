from fun.vac.algorithms.sorting import MergeSort
from fun.vac.util import SequenceBuilder
from fun.vac.util import Sequences
from fun.vac.util import TestRunner

def testMergeSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    MergeSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    if not Sequences.isSorted(arr):
        return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testMergeSort())
