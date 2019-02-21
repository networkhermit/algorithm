package SequenceBuilder

import (
    "fun/vac/util/RandomFactory"
    "fun/vac/util/Sequences"
)

func PackIncreasing(arr []int) {
    RandomFactory.Launch()
    arr[0] = RandomFactory.IntegerN(1, 4)
    for i, length := 1, len(arr); i < length; i++ {
        arr[i] = arr[i - 1] + RandomFactory.IntegerN(1, 4)
    }
}

func PackRandom(arr []int) {
    RandomFactory.Launch()
    for i, _ := range arr {
        arr[i] = RandomFactory.GenerateInteger()
    }
}

func PackDecreasing(arr []int) {
    PackIncreasing(arr)
    Sequences.Reverse(arr)
}
