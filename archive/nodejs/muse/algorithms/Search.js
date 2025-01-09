'use strict';

const BinarySearch = require('muse/algorithms/search/BinarySearch');
const LinearSearch = require('muse/algorithms/search/LinearSearch');

exports.binarySearch = (arr, key) => {
  return BinarySearch.find(arr, key);
};

exports.linearSearch = (arr, key) => {
  return LinearSearch.find(arr, key);
};
