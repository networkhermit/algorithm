package quick_sort

func Partition(arr []int, lo, hi int) int {
	pivot := arr[lo]

	left := lo
	right := hi - 1

	for left != right {
		for i := right; i > left; i-- {
			if arr[i] < pivot {
				arr[left] = arr[i]
				arr[i] = pivot
				break
			}
			right--
		}
		for i := left; i < right; i++ {
			if arr[i] > pivot {
				arr[right] = arr[i]
				arr[i] = pivot
				break
			}
			left++
		}
	}

	return left
}

func Sort(arr []int) {
	SortWithRange(arr, 0, len(arr))
}

func SortWithRange(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := Partition(arr, lo, hi)

	SortWithRange(arr, lo, p)
	SortWithRange(arr, p+1, hi)
}
