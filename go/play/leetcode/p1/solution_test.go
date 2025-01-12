package p1

import (
	"slices"
	"testing"
)

func Test(t *testing.T) {
	sample := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range sample {
		actual := twoSum(tc.nums, tc.target)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("%s(%v, %d) returned %v, expect %v", t.Name(), tc.nums, tc.target, actual, tc.expected)
		}
	}
}

func TestBruteForce(t *testing.T) {
	sample := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range sample {
		actual := twoSumBruteForce(tc.nums, tc.target)
		if !slices.Equal(actual, tc.expected) {
			t.Errorf("%s(%v, %d) returned %v, expect %v", t.Name(), tc.nums, tc.target, actual, tc.expected)
		}
	}
}
