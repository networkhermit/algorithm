from muse.algorithms.sorting import (
    bubble_sort,
    insertion_sort,
    merge_sort,
    quick_sort,
    selection_sort,
)


def bubblesort(arr: list[int]) -> None:
    bubble_sort.sort(arr)


def insertionsort(arr: list[int]) -> None:
    insertion_sort.sort(arr)


def mergesort(arr: list[int]) -> None:
    merge_sort.sort(arr)


def quicksort(arr: list[int]) -> None:
    quick_sort.sort(arr)


def selectionsort(arr: list[int]) -> None:
    selection_sort.sort(arr)
