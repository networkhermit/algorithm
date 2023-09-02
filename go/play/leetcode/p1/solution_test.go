package p1

import (
	"slices"
	"testing"
)

func Test(t *testing.T) {
	var sample = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range sample {
		actual := twoSum(tt.nums, tt.target)
		if !slices.Equal(actual, tt.expected) {
			t.Errorf("%s(%v, %d) returned %v, expected %v", t.Name(), tt.nums, tt.target, actual, tt.expected)
		}
	}
}

func TestBruteForce(t *testing.T) {
	var sample = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range sample {
		actual := twoSumBruteForce(tt.nums, tt.target)
		if !slices.Equal(actual, tt.expected) {
			t.Errorf("%s(%v, %d) returned %v, expected %v", t.Name(), tt.nums, tt.target, actual, tt.expected)
		}
	}
}
