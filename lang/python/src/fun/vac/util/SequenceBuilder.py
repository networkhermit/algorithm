from fun.vac.util import RandomFactory
from fun.vac.util import Sequences

def packIncreasing(arr: list) -> None:
    RandomFactory.launch()
    arr[0] = RandomFactory.integerN(1, 4)
    for i in range(1, len(arr)):
        arr[i] = arr[i - 1] + RandomFactory.integerN(1, 4)

def packRandom(arr: list) -> None:
    RandomFactory.launch()
    for i in range(len(arr)):
        arr[i] = RandomFactory.generateInteger()

def packDecreasing(arr: list) -> None:
    packIncreasing(arr)
    Sequences.reverse(arr)
