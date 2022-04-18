'use strict'

exports.find = (arr, key) => {
  for (let i = 0, length = arr.length; i < length; i++) {
    if (arr[i] === key) {
      return i
    }
  }

  return arr.length
}
