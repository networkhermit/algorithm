package BinarySearch

func Find(arr []int, key int) int {
    lo := 0
    hi := len(arr)

    var mid int

    for lo < hi {
        mid = lo + ((hi - lo) >> 1)
        if key < arr[mid] {
            hi = mid
        } else if key > arr[mid] {
            lo = mid + 1
        } else {
            return mid
        }
    }

    return len(arr)
}
