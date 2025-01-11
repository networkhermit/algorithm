package selection_sort

func Sort(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := range len(arr) - 1 {
		iMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
		}
		if iMin != i {
			arr[i], arr[iMin] = arr[iMin], arr[i]
		}
	}
}
