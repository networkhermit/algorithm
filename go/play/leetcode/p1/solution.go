package p1

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for j, n := range nums {
		m := target - n
		if i, ok := hashmap[m]; ok {
			return []int{i, j}
		}
		hashmap[n] = j
	}

	return []int{}
}

func twoSumBruteForce(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
