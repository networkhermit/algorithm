package BinarySearch

func Find(arr []int, key int) int {
	lo := 0
	hi := len(arr)

	var mid int

	for lo < hi {
		mid = lo + ((hi - lo) >> 1)
		switch {
		case key < arr[mid]:
			hi = mid
		case key > arr[mid]:
			lo = mid + 1
		default:
			return mid
		}
	}

	return len(arr)
}
