export const partition = (arr, lo, hi) => {
  if (lo === hi) {
    return;
  }

  const pivot = arr[lo];

  let left = lo;
  let right = hi - 1;

  while (left !== right) {
    for (let i = right; i > left; i--) {
      if (arr[i] < pivot) {
        arr[left] = arr[i];
        arr[i] = pivot;
        break;
      }
      right--;
    }
    for (let i = left; i < right; i++) {
      if (arr[i] > pivot) {
        arr[right] = arr[i];
        arr[i] = pivot;
        break;
      }
      left++;
    }
  }

  partition(arr, lo, left);
  partition(arr, left + 1, hi);
};

export const sort = (arr) => {
  partition(arr, 0, arr.length);
};
