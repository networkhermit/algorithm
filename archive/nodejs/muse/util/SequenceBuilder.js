import * as RandomFactory from './RandomFactory.js';
import * as Sequences from './Sequences.js';

export const packIncreasing = (arr) => {
  if (arr.length === 0) {
    return;
  }
  arr[0] = RandomFactory.genIntN(1, 3);
  for (let i = 1, length = arr.length; i < length; i++) {
    arr[i] = arr[i - 1] + RandomFactory.genIntN(1, 3);
  }
};

export const packRandom = (arr) => {
  for (let i = 0, length = arr.length; i < length; i++) {
    arr[i] = RandomFactory.genInt();
  }
};

export const packDecreasing = (arr) => {
  this.packIncreasing(arr);
  Sequences.reverse(arr);
};
