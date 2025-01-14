package shell_sort

func Sort(arr []int) {
	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}
	for _, gap := range gaps {
		for i := gap; i < len(arr); i++ {
			target := arr[i]
			cursor := i
			for ; cursor >= gap; cursor -= gap {
				if arr[cursor-gap] <= target {
					break
				}
				arr[cursor] = arr[cursor-gap]
			}
			arr[cursor] = target
		}
	}
}
