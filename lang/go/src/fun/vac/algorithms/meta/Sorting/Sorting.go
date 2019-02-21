package Sorting

import (
    bubble_sort    "fun/vac/algorithms/sorting/BubbleSort"
    insertion_sort "fun/vac/algorithms/sorting/InsertionSort"
    merge_sort     "fun/vac/algorithms/sorting/MergeSort"
    quick_sort     "fun/vac/algorithms/sorting/QuickSort"
    selection_sort "fun/vac/algorithms/sorting/SelectionSort"
)

func BubbleSort(arr []int) {
    bubble_sort.Sort(arr)
}

func InsertionSort(arr []int) {
    insertion_sort.Sort(arr)
}

func MergeSort(arr []int) {
    merge_sort.Sort(arr)
}

func QuickSort(arr []int) {
    quick_sort.Sort(arr)
}

func SelectionSort(arr []int) {
    selection_sort.Sort(arr)
}
