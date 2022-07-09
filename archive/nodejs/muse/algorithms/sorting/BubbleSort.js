'use strict'

exports.sort = (arr) => {
  let temp = null

  let margin = 0
  let unsorted = arr.length

  while (unsorted > 1) {
    margin = 0
    for (let i = 1; i < unsorted; i++) {
      if (arr[i - 1] > arr[i]) {
        temp = arr[i - 1]
        arr[i - 1] = arr[i]
        arr[i] = temp
        margin = i
      }
    }
    unsorted = margin
  }
}
