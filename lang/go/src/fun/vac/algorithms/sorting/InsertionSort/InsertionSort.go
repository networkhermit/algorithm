package InsertionSort

func Sort(arr []int) {
    var target int

    var cursor int

    for i, length := 1, len(arr); i < length; i++ {
        target = arr[i]
        for cursor = i - 1; cursor >= 0; cursor-- {
            if arr[cursor] > target {
                arr[cursor + 1] = arr[cursor]
            } else {
                break
            }
        }
        arr[cursor + 1] = target
    }
}
