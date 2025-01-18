export const sort = (arr) => {
  let target = null;

  let cursor = 0;

  for (let i = 1, length = arr.length; i < length; i++) {
    target = arr[i];
    for (cursor = i; cursor > 0; cursor--) {
      if (arr[cursor - 1] <= target) {
        break;
      }
      arr[cursor] = arr[cursor - 1];
    }
    arr[cursor] = target;
  }
};
