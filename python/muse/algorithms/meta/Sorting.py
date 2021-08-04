from muse.algorithms.sorting import BubbleSort
from muse.algorithms.sorting import InsertionSort
from muse.algorithms.sorting import MergeSort
from muse.algorithms.sorting import QuickSort
from muse.algorithms.sorting import SelectionSort

def bubbleSort(arr: list) -> None:
    BubbleSort.sort(arr)

def insertionSort(arr: list) -> None:
    InsertionSort.sort(arr)

def mergeSort(arr: list) -> None:
    MergeSort.sort(arr)

def quickSort(arr: list) -> None:
    QuickSort.sort(arr)

def selectionSort(arr: list) -> None:
    SelectionSort.sort(arr)
