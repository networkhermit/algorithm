package SelectionSort

import (
    "testing"

    "muse/util/SequenceBuilder"
    "muse/util/Sequences"
)

func TestSelectionSort(t *testing.T) {
    size := 32768

    arr := make([]int, size)
    SequenceBuilder.PackRandom(arr)

    checksum := Sequences.ParityChecksum(arr)

    Sort(arr)

    if Sequences.ParityChecksum(arr) != checksum {
        t.FailNow()
    }

    if !Sequences.IsSorted(arr) {
        t.FailNow()
    }
}
