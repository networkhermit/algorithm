package Search

import (
    "testing"

    "muse/util/SequenceBuilder"
)

func TestBinarySearch(t *testing.T) {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if BinarySearch(arr, -1) != size {
        t.FailNow()
    }

    if BinarySearch(arr, 2147483647) != size {
        t.FailNow()
    }

    for i, v := range arr {
        if BinarySearch(arr, v) != i {
            t.FailNow()
        }
    }
}

func TestLinearSearch(t *testing.T) {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackIncreasing(arr)

    if LinearSearch(arr, -1) != size {
        t.FailNow()
    }

    if LinearSearch(arr, 2147483647) != size {
        t.FailNow()
    }

    for i, v := range arr {
        if LinearSearch(arr, v) != i {
            t.FailNow()
        }
    }
}
