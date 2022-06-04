package sequence_builder

import (
	"muse/util/random_factory"
	"muse/util/sequences"
)

func PackIncreasing(arr []int) {
	arr[0] = random_factory.GenIntN(1, 3)
	for i, length := 1, len(arr); i < length; i++ {
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
