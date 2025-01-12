package tests

import (
	"testing"

	"muse/util/sequence_builder"
)

var Derive = func(find func([]int, int) int, size int, pack func([]int)) func(t *testing.T) {
	return func(t *testing.T) {
		arr := make([]int, size)
		pack(arr)

		sentinel := []int{-1, 2_147_483_647}

		for _, v := range sentinel {
			if actual := find(arr, v); actual != size {
				t.Errorf("%s(arr, %d) returned %d, expect %d", t.Name(), v, actual, size)
			}
		}

		for i, v := range arr {
			if actual := find(arr, v); actual != i {
				t.Errorf("%s(arr, %d) returned %d, expect %d", t.Name(), i, actual, i)
			}
		}
	}
}

var DeriveEmpty = func(find func([]int, int) int) func(t *testing.T) {
	return Derive(find, 0, sequence_builder.PackIdentical)
}

var DeriveIncreasing = func(find func([]int, int) int, size int) func(t *testing.T) {
	return Derive(find, size, sequence_builder.PackIncreasing)
}
