package sequences

import (
	"fmt"
	"sort"

	"muse/util/random_factory"
)

func Inspect(arr []int) {
	fmt.Println("[")
	for i, v := range arr {
		fmt.Printf("\t#%04X  ->  %d\n", i, v)
	}
	fmt.Println("]")
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}

func ParityChecksum(arr []int) int {
	checksum := 0

	for _, v := range arr {
		checksum ^= v
	}

	return checksum
}

func Reverse(arr []int) {
	var k int

	for i, bound := 0, len(arr)>>1; i < bound; i++ {
		k = len(arr) - i - 1
		arr[i], arr[k] = arr[k], arr[i]
	}
}

func Shuffle(arr []int) {
	var k int

	for i := range len(arr) {
		k = random_factory.GenIntN(i, len(arr)-1)
		arr[i], arr[k] = arr[k], arr[i]
	}
}

func Sort(arr []int) {
	sort.Ints(arr)
}
