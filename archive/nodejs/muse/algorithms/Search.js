import * as BinarySearch from './search/BinarySearch.js';
import * as LinearSearch from './search/LinearSearch.js';

export const binarySearch = (arr, key) => {
  return BinarySearch.find(arr, key);
};

export const linearSearch = (arr, key) => {
  return LinearSearch.find(arr, key);
};
