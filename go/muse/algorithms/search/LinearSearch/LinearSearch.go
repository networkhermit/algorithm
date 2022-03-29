package LinearSearch

func Find(arr []int, key int) int {
	for i, v := range arr {
		if v == key {
			return i
		}
	}

	return len(arr)
}
