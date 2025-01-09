'use strict';

exports.merge = (arr, lo, mid, hi) => {
  if (lo === mid) {
    return;
  }

  this.merge(arr, lo, (lo + mid) >>> 1, mid);
  this.merge(arr, mid, (mid + hi) >>> 1, hi);

  let m = lo;
  let n = mid;

  const sorted = new Array(hi - lo);

  for (let i = 0, length = sorted.length; i < length; i++) {
    if (m !== mid && (n === hi || arr[m] < arr[n])) {
      sorted[i] = arr[m];
      m++;
    } else {
      sorted[i] = arr[n];
      n++;
    }
  }

  let cursor = 0;
  for (let i = lo; i < hi; i++) {
    arr[i] = sorted[cursor];
    cursor++;
  }
};

exports.sort = (arr) => {
  this.merge(arr, 0, arr.length >>> 1, arr.length);
};
