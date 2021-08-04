"use strict"

exports.partition = (arr, lo, hi) => {
    if (lo === hi) {
        return
    }

    let pivot = arr[lo]

    let left = lo
    let right = hi - 1

    while (left !== right) {
        for (let i = right; i > left; i--) {
            if (arr[i] < pivot) {
                arr[left] = arr[i]
                arr[i] = pivot
                break
            }
            right--
        }
        for (let i = left; i < right; i++) {
            if (arr[i] > pivot) {
                arr[right] = arr[i]
                arr[i] = pivot
                break
            }
            left++
        }
    }

    this.partition(arr, lo, left)
    this.partition(arr, left + 1, hi)
}

exports.sort = (arr) => {
    this.partition(arr, 0, arr.length)
}
