package main

import "fmt"

func twoSum(nums []int, target int) []int {
    indices := make([]int, 2)

    for i, length := 0, len(nums); i < length; i++ {
        for j := i + 1; j < length; j++ {
            if nums[i] + nums[j] == target {
                indices[0] = i
                indices[1] = j
                return indices
            }
        }
    }

    return indices
}

func main() {
    nums := []int{2, 7, 11, 15}

    target := 9

    indices := twoSum(nums, target)

    fmt.Println(indices[0] == 0 && indices[1] == 1)
}
