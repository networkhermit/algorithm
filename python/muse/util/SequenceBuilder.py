from muse.util import RandomFactory, Sequences


def packIncreasing(arr: list) -> None:
    arr[0] = RandomFactory.genIntN(1, 3)
    for i in range(1, len(arr)):
        arr[i] = arr[i - 1] + RandomFactory.genIntN(1, 3)


def packRandom(arr: list) -> None:
    for i, _ in enumerate(arr):
        arr[i] = RandomFactory.genInt()


def packDecreasing(arr: list) -> None:
    packIncreasing(arr)
    Sequences.reverse(arr)
