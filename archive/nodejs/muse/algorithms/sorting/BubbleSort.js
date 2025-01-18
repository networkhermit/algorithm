export const sort = (arr) => {
  let margin = 0;
  let unsorted = arr.length;

  while (unsorted > 1) {
    margin = 0;
    for (let i = 1; i < unsorted; i++) {
      if (arr[i - 1] > arr[i]) {
        [arr[i - 1], arr[i]] = [arr[i], arr[i - 1]];
        margin = i;
      }
    }
    unsorted = margin;
  }
};
