package quick_sort

import "math/rand"

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
	if lo+1 >= hi {
		return
	}

	p := Partition(arr, lo, hi)

	SortWithRange(arr, lo, p+1)
	SortWithRange(arr, p+1, hi)
}

func PartitionThreeWay(arr []int, lo, hi int) (int, int) {
	i := lo - 1
	j := hi - 1

	p := lo - 1
	q := hi - 1

	pivot := arr[hi-1]

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if pivot >= arr[j] {
				break
			}
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		if arr[i] == pivot {
			p++
			arr[p], arr[i] = arr[i], arr[p]
		}
		if pivot == arr[j] {
			q--
			arr[j], arr[q] = arr[q], arr[j]
		}
	}

	arr[i], arr[hi-1] = arr[hi-1], arr[i]

	j = i - 1
	i++

	for k := lo; k < p; k++ {
		arr[k], arr[j] = arr[j], arr[k]
		j--
	}
	for k := hi - 2; k > q; k-- {
		arr[i], arr[k] = arr[k], arr[i]
		i++
	}

	return i, j
}

func SortThreeWayPartitionWithShuffle(arr []int) {
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	SortWithRangeThreeWayPartition(arr, 0, len(arr))
}

func SortThreeWayPartition(arr []int) {
	SortWithRangeThreeWayPartition(arr, 0, len(arr))
}

func SortWithRangeThreeWayPartition(arr []int, lo, hi int) {
	if lo+1 >= hi {
		return
	}

	i, j := PartitionThreeWay(arr, lo, hi)

	SortWithRangeThreeWayPartition(arr, lo, j+1)
	SortWithRangeThreeWayPartition(arr, i, hi)
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
