package Search

import (
	binary_search "muse/algorithms/search/BinarySearch"
	linear_search "muse/algorithms/search/LinearSearch"
)

func BinarySearch(arr []int, key int) int {
	return binary_search.Find(arr, key)
}

func LinearSearch(arr []int, key int) int {
	return linear_search.Find(arr, key)
}
