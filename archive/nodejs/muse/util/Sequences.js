import * as QuickSort from '../algorithms/sorting/QuickSort.js';
import * as RandomFactory from './RandomFactory.js';

export const inspect = (arr) => {
  console.log('[');
  for (let i = 0, length = arr.length; i < length; i++) {
    process.stdout.write(
      `\t#${i.toString(16).toUpperCase().padStart(4, '0')}  -->  ${arr[i]}\n`,
    );
  }
  console.log(']');
};

export const isSorted = (arr) => {
  for (let i = 1, length = arr.length; i < length; i++) {
    if (arr[i] < arr[i - 1]) {
      return false;
    }
  }

  return true;
};

export const parityChecksum = (arr) => {
  let checksum = 0;

  for (const v of arr) {
    checksum ^= v;
  }

  return checksum;
};

export const reverse = (arr) => {
  let k;

  for (
    let i = 0, bound = Math.floor(arr.length >>> 1), length = arr.length;
    i < bound;
    i++
  ) {
    k = length - i - 1;
    [arr[i], arr[k]] = [arr[k], arr[i]];
  }
};

export const shuffle = (arr) => {
  let k;

  for (let i = 0, length = arr.length; i < length; i++) {
    k = RandomFactory.genIntN(i, length - 1);
    [arr[i], arr[k]] = [arr[k], arr[i]];
  }
};

export const sort = (arr) => {
  QuickSort.sort(arr);
};
