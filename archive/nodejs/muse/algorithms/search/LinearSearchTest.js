'use strict';

const LinearSearch = require('muse/algorithms/search/LinearSearch');
const SequenceBuilder = require('muse/util/SequenceBuilder');
const TestRunner = require('muse/util/TestRunner');

const testLinearSearch = () => {
  const size = 32768;

  const arr = new Array(size);
  SequenceBuilder.packIncreasing(arr);

  if (LinearSearch.find(arr, -1) !== size) {
    return false;
  }

  if (LinearSearch.find(arr, 2_147_483_647) !== size) {
    return false;
  }

  for (let i = 0; i < size; i++) {
    if (LinearSearch.find(arr, arr[i]) !== i) {
      return false;
    }
  }

  return true;
};

if (module === require.main) {
  TestRunner.pick(testLinearSearch);
}
