package Sorting

import (
	bubble_sort "muse/algorithms/sorting/BubbleSort"
	insertion_sort "muse/algorithms/sorting/InsertionSort"
	merge_sort "muse/algorithms/sorting/MergeSort"
	quick_sort "muse/algorithms/sorting/QuickSort"
	selection_sort "muse/algorithms/sorting/SelectionSort"
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
