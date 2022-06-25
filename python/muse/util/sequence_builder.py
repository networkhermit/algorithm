from muse.util import random_factory, sequences


def pack_increasing(arr: list) -> None:
    arr[0] = random_factory.gen_int_n(1, 3)
    for i in range(1, len(arr)):
        arr[i] = arr[i - 1] + random_factory.gen_int_n(1, 3)


def pack_random(arr: list) -> None:
    for i, _ in enumerate(arr):
        arr[i] = random_factory.gen_int()


def pack_decreasing(arr: list) -> None:
    pack_increasing(arr)
    sequences.reverse(arr)
