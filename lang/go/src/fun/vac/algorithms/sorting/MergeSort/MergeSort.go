package MergeSort

func Merge(arr []int, lo int, mid int, hi int) {
    if lo == mid {
        return
    }

    Merge(arr, lo, (lo + mid) >> 1, mid)
    Merge(arr, mid, (mid + hi) >> 1, hi)

    m := lo
    n := mid

    sorted := make([]int, hi - lo)

    for i, length := 0, len(sorted); i < length; i++ {
        if m != mid && (n == hi || arr[m] < arr[n]) {
            sorted[i] = arr[m]
            m += 1
        } else {
            sorted[i] = arr[n]
            n += 1
        }
    }

    cursor := 0
    for i := lo; i < hi; i++ {
        arr[i] = sorted[cursor]
        cursor += 1
    }
}

func Sort(arr []int) {
    Merge(arr, 0, len(arr) >> 1, len(arr))
}
