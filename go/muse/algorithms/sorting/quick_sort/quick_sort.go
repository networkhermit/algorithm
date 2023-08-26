package quick_sort

func Partition(arr []int, lo, hi int) int {
	pivot := arr[lo+((hi-1-lo)>>1)]

	i := lo - 1
	j := hi

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Sort(arr []int) {
	SortWithRange(arr, 0, len(arr))
}

func SortWithRange(arr []int, lo, hi int) {
	if lo >= hi-1 {
		return
	}

	p := Partition(arr, lo, hi)

	SortWithRange(arr, lo, p+1)
	SortWithRange(arr, p+1, hi)
}

func PartitionInefficient(arr []int, lo, hi int) int {
	pivot := arr[lo]

	i := lo
	j := hi - 1

	for i != j {
		for j > i {
			if arr[j] < pivot {
				arr[i], arr[j] = arr[j], pivot
				break
			}
			j--
		}
		for i < j {
			if arr[i] > pivot {
				arr[j], arr[i] = arr[i], pivot
				break
			}
			i++
		}
	}

	return i
}

func SortInefficient(arr []int) {
	SortWithRangeInefficient(arr, 0, len(arr))
}

func SortWithRangeInefficient(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	p := PartitionInefficient(arr, lo, hi)

	SortWithRangeInefficient(arr, lo, p)
	SortWithRangeInefficient(arr, p+1, hi)
}
