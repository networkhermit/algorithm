package QuickSort

func Partition(arr []int, lo int, hi int) {
    if lo == hi {
        return
    }

    pivot := arr[lo]

    left, right := lo, hi - 1

    for left != right {
        for i := right; i > left; i-- {
            if arr[i] < pivot {
                arr[left] = arr[i]
                arr[i] = pivot
                break
            }
            right -= 1
        }
        for i := left; i < right; i++ {
            if arr[i] > pivot {
                arr[right] = arr[i]
                arr[i] = pivot
                break
            }
            left += 1
        }
    }

    Partition(arr, lo, left)
    Partition(arr, left + 1, hi)
}

func Sort(arr []int) {
    Partition(arr, 0, len(arr))
}
