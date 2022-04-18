'use strict'

exports.find = (arr, key) => {
  let lo = 0
  let hi = arr.length

  let mid = 0

  while (lo < hi) {
    mid = lo + ((hi - lo) >>> 1)
    switch (true) {
      case key < arr[mid]:
        hi = mid
        break
      case key > arr[mid]:
        lo = mid + 1
        break
      default:
        return mid
    }
  }

  return arr.length
}
