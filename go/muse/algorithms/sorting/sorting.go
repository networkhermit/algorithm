package sorting

import (
	"muse/algorithms/sorting/bubble_sort"
	"muse/algorithms/sorting/insertion_sort"
	"muse/algorithms/sorting/merge_sort"
	"muse/algorithms/sorting/quick_sort"
	"muse/algorithms/sorting/selection_sort"
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
