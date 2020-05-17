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

    return Sequences.isSorted(arr)

def main() -> None:
    TestRunner.parseTest(testQuickSort())

if __name__ == "__main__":
    main()
