"use strict"

exports.find = (arr, key) => {
    let lo = 0
    let hi = arr.length

    let mid = 0

    while (lo < hi) {
        mid = lo + ((hi - lo) >>> 1)
        if (key < arr[mid]) {
            hi = mid
        } else if (key > arr[mid]) {
            lo = mid + 1
        } else {
            return mid
        }
    }

    return arr.length
}
