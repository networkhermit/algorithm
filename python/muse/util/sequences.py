from muse.algorithms.sorting import quick_sort
from muse.util import random_factory


def inspect(arr: list) -> None:
    print("[")
    for i, v in enumerate(arr):
        print("\t#%04X  -->  %d" % (i, v))
    print("]")


def is_sorted(arr: list) -> bool:
    for i in range(1, len(arr)):
        if arr[i] < arr[i - 1]:
            return False

    return True


def parity_checksum(arr: list) -> int:
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
        k = random_factory.gen_int_n(i, length - 1)
        arr[i], arr[k] = arr[k], arr[i]


def sort(arr: list) -> None:
    quick_sort.sort(arr)
