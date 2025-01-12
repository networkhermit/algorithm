from muse.algorithms.sorting import quick_sort
from muse.util import random_factory


def inspect(arr: list[int]) -> None:
    print("[")
    for i, v in enumerate(arr):
        print("\t#%04X  -->  %d" % (i, v))
    print("]")


def is_sorted(arr: list[int]) -> bool:
    for i in range(1, len(arr)):
        if arr[i] < arr[i - 1]:
            return False

    return True


def parity_checksum(arr: list[int]) -> int:
    checksum = 0

    for v in arr:
        checksum ^= v

    return checksum


def reverse(arr: list[int]) -> None:
    for i in range(len(arr) >> 1):
        k = len(arr) - i - 1
        arr[i], arr[k] = arr[k], arr[i]


def shuffle(arr: list[int]) -> None:
    for i in range(len(arr)):
        k = random_factory.gen_int_n(i, len(arr) - 1)
        arr[i], arr[k] = arr[k], arr[i]


def sort(arr: list[int]) -> None:
    quick_sort.sort(arr)
