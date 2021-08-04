from muse.util import RandomFactory
from muse.util import Sequences

def packIncreasing(arr: list) -> None:
    RandomFactory.seed()
    arr[0] = RandomFactory.integerN(1, 3)
    for i in range(1, len(arr)):
        arr[i] = arr[i - 1] + RandomFactory.integerN(1, 3)

def packRandom(arr: list) -> None:
    RandomFactory.seed()
    for i, _ in enumerate(arr):
        arr[i] = RandomFactory.generateInteger()

def packDecreasing(arr: list) -> None:
    packIncreasing(arr)
    Sequences.reverse(arr)
