'use strict'

exports.sort = (arr) => {
  if (arr.length === 0) {
    return
  }

  let iMin = 0

  for (let i = 0, bound = arr.length - 1; i < bound; i++) {
    iMin = i
    for (let j = i + 1; j <= bound; j++) {
      if (arr[j] < arr[iMin]) {
        iMin = j
      }
    }
    if (iMin !== i) {
      [arr[i], arr[iMin]] = [arr[iMin], arr[i]]
    }
  }
}
