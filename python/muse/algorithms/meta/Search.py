from muse.algorithms.search import BinarySearch, LinearSearch


def binarySearch(arr: list, key: object) -> int:
    return BinarySearch.find(arr, key)


def linearSearch(arr: list, key: object) -> int:
    return LinearSearch.find(arr, key)
