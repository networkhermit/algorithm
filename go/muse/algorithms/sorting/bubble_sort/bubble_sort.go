package bubble_sort

func Sort(arr []int) {
	unsorted := len(arr)
	for unsorted > 1 {
		margin := 0
		for i := 1; i < unsorted; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				margin = i
			}
		}
		unsorted = margin
	}
}
