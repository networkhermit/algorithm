from fun.vac.algorithms.sorting import InsertionSort
from fun.vac.util import SequenceBuilder
from fun.vac.util import Sequences
from fun.vac.util import TestRunner

def testInsertionSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    InsertionSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def main() -> None:
    TestRunner.parseTest(testInsertionSort())

if __name__ == "__main__":
    main()
