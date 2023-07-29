package quick_sort

func Partition(arr []int, lo, hi int) {
	if lo == hi {
		return
	}

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

	Partition(arr, lo, left)
	Partition(arr, left+1, hi)
}

func Sort(arr []int) {
	Partition(arr, 0, len(arr))
}
