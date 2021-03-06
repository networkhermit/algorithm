from muse.algorithms.sorting import BubbleSort
from muse.util import SequenceBuilder
from muse.util import Sequences
from muse.util import TestRunner

def testBubbleSort() -> bool:
    size = 32768

    arr = [0] * size
    SequenceBuilder.packRandom(arr)

    checksum = Sequences.parityChecksum(arr)

    BubbleSort.sort(arr)

    if Sequences.parityChecksum(arr) != checksum:
        return False

    return Sequences.isSorted(arr)

def main() -> None:
    TestRunner.parseTest(testBubbleSort())

if __name__ == "__main__":
    main()
