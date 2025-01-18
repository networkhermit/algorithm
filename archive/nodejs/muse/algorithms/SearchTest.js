import * as SequenceBuilder from '../util/SequenceBuilder.js';
import * as TestRunner from '../util/TestRunner.js';
import * as Search from './Search.js';

const testBinarySearch = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packIncreasing(arr);

  if (Search.binarySearch(arr, -1) !== size) {
    return false;
  }

  if (Search.binarySearch(arr, 2_147_483_647) !== size) {
    return false;
  }

  for (let i = 0; i < size; i++) {
    if (Search.binarySearch(arr, arr[i]) !== i) {
      return false;
    }
  }

  return true;
};

const testLinearSearch = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packIncreasing(arr);

  if (Search.linearSearch(arr, -1) !== size) {
    return false;
  }

  if (Search.linearSearch(arr, 2_147_483_647) !== size) {
    return false;
  }

  for (let i = 0; i < size; i++) {
    if (Search.linearSearch(arr, arr[i]) !== i) {
      return false;
    }
  }

  return true;
};

const main = () => {
  TestRunner.pick(testBinarySearch);

  TestRunner.pick(testLinearSearch);
};

main();
