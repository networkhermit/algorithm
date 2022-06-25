from muse.algorithms.search import binary_search, linear_search


def binarysearch(arr: list, key: object) -> int:
    return binary_search.find(arr, key)


def linearsearch(arr: list, key: object) -> int:
    return linear_search.find(arr, key)
