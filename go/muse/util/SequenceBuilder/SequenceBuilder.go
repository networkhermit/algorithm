package SequenceBuilder

import (
    "muse/util/RandomFactory"
    "muse/util/Sequences"
)

func PackIncreasing(arr []int) {
    RandomFactory.Seed()
    arr[0] = RandomFactory.GenIntN(1, 3)
    for i, length := 1, len(arr); i < length; i++ {
        arr[i] = arr[i - 1] + RandomFactory.GenIntN(1, 3)
    }
}

func PackRandom(arr []int) {
    RandomFactory.Seed()
    for i := range arr {
        arr[i] = RandomFactory.GenInt()
    }
}

func PackDecreasing(arr []int) {
    PackIncreasing(arr)
    Sequences.Reverse(arr)
}
