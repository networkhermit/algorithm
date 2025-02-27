package merge_sort

func Merge(arr []int, lo, mid, hi int) {
	if lo == mid {
		return
	}

	Merge(arr, lo, lo+((mid-lo)>>1), mid)
	Merge(arr, mid, mid+((hi-mid)>>1), hi)

	m := lo
	n := mid

	sorted := make([]int, hi-lo)

	for i := range sorted {
		if m != mid && (n == hi || arr[m] < arr[n]) {
			sorted[i] = arr[m]
			m++
		} else {
			sorted[i] = arr[n]
			n++
		}
	}

	cursor := 0
	for i := lo; i < hi; i++ {
		arr[i] = sorted[cursor]
		cursor++
	}
}

func Sort(arr []int) {
	Merge(arr, 0, len(arr)>>1, len(arr))
}
