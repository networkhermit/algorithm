import * as BubbleSort from './sorting/BubbleSort.js';
import * as InsertionSort from './sorting/InsertionSort.js';
import * as MergeSort from './sorting/MergeSort.js';
import * as QuickSort from './sorting/QuickSort.js';
import * as SelectionSort from './sorting/SelectionSort.js';

export const bubbleSort = (arr) => {
  BubbleSort.sort(arr);
};

export const insertionSort = (arr) => {
  InsertionSort.sort(arr);
};

export const mergeSort = (arr) => {
  MergeSort.sort(arr);
};

export const quickSort = (arr) => {
  QuickSort.sort(arr);
};

export const selectionSort = (arr) => {
  SelectionSort.sort(arr);
};
