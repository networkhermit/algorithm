package sequence_builder

import (
	"muse/util/random_factory"
	"muse/util/sequences"
)

func PackIdentical(arr []int) {
	n := random_factory.GenInt()
	for i := range arr {
		arr[i] = n
	}
}

func PackIncreasing(arr []int) {
	if len(arr) == 0 {
		return
	}
	arr[0] = random_factory.GenIntN(1, 3)
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + random_factory.GenIntN(1, 3)
	}
}

func PackRandom(arr []int) {
	for i := range arr {
		arr[i] = random_factory.GenInt()
	}
}

func PackDecreasing(arr []int) {
	PackIncreasing(arr)
	sequences.Reverse(arr)
}
