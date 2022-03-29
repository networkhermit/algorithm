from muse.algorithms.sorting import QuickSort
from muse.util import RandomFactory


def inspect(arr: list) -> None:
    print("[")
    for i, v in enumerate(arr):
        print("\t#%04X  -->  %d" % (i, v))
    print("]")


def isSorted(arr: list) -> bool:
    for i in range(1, len(arr)):
        if arr[i] < arr[i - 1]:
            return False

    return True


def parityChecksum(arr: list) -> int:
    checksum = 0

    for v in arr:
        checksum ^= v

    return checksum


def reverse(arr: list) -> None:
    k = 0

    length = len(arr)
    for i in range(len(arr) >> 1):
        k = length - i - 1
        arr[i], arr[k] = arr[k], arr[i]


def shuffle(arr: list) -> None:
    k = 0
    length = len(arr)

    for i in range(length):
        k = RandomFactory.genIntN(i, length)
        arr[i], arr[k] = arr[k], arr[i]


def sort(arr: list) -> None:
    QuickSort.sort(arr)
