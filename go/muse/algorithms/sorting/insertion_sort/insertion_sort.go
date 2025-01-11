package insertion_sort

func Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		target := arr[i]
		cursor := i
		for ; cursor > 0; cursor-- {
			if arr[cursor-1] <= target {
				break
			}
			arr[cursor] = arr[cursor-1]
		}
		arr[cursor] = target
	}
}
