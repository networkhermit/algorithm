from muse.algorithms.sorting import (
    bubble_sort,
    insertion_sort,
    merge_sort,
    quick_sort,
    selection_sort,
)


def bubblesort(arr: list) -> None:
    bubble_sort.sort(arr)


def insertionsort(arr: list) -> None:
    insertion_sort.sort(arr)


def mergesort(arr: list) -> None:
    merge_sort.sort(arr)


def quicksort(arr: list) -> None:
    quick_sort.sort(arr)


def selectionsort(arr: list) -> None:
    selection_sort.sort(arr)
