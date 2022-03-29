package SelectionSort

func Sort(arr []int) {
	if len(arr) == 0 {
		return
	}

	var iMin int

	for i, bound := 0, len(arr)-1; i < bound; i++ {
		iMin = i
		for j := i + 1; j <= bound; j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
		}
		if iMin != i {
			arr[i], arr[iMin] = arr[iMin], arr[i]
		}
	}
}
