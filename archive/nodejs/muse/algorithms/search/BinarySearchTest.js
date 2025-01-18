import * as SequenceBuilder from '../../util/SequenceBuilder.js';
import * as TestRunner from '../../util/TestRunner.js';
import * as BinarySearch from './BinarySearch.js';

const testBinarySearch = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packIncreasing(arr);

  if (BinarySearch.find(arr, -1) !== size) {
    return false;
  }

  if (BinarySearch.find(arr, 2_147_483_647) !== size) {
    return false;
  }

  for (let i = 0; i < size; i++) {
    if (BinarySearch.find(arr, arr[i]) !== i) {
      return false;
    }
  }

  return true;
};

const main = () => {
  TestRunner.pick(testBinarySearch);
};

main();
