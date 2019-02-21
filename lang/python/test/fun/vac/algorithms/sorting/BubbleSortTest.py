from fun.vac.algorithms.sorting import BubbleSort
from fun.vac.util import SequenceBuilder
from fun.vac.util import Sequences
from fun.vac.util import TestRunner

def testBubbleSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    BubbleSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    if not Sequences.isSorted(arr):
        return False

    return True

if __name__ == "__main__":
    TestRunner.parseTest(testBubbleSort())
