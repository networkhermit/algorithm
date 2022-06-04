package search

import (
	"muse/algorithms/search/binary_search"
	"muse/algorithms/search/linear_search"
)

func BinarySearch(arr []int, key int) int {
	return binary_search.Find(arr, key)
}

func LinearSearch(arr []int, key int) int {
	return linear_search.Find(arr, key)
}
