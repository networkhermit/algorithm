#ifndef MUSE_ALGORITHMS_META_SORTING_HPP
#define MUSE_ALGORITHMS_META_SORTING_HPP

#include <muse/algorithms/sorting/BubbleSort.hpp>
#include <muse/algorithms/sorting/InsertionSort.hpp>
#include <muse/algorithms/sorting/MergeSort.hpp>
#include <muse/algorithms/sorting/QuickSort.hpp>
#include <muse/algorithms/sorting/SelectionSort.hpp>

namespace Sorting {

template <typename T> void bubbleSort(std::vector<T> &arr) {
  BubbleSort::sort(arr);
}

template <typename T> void insertionSort(std::vector<T> &arr) {
  InsertionSort::sort(arr);
}

template <typename T> void mergeSort(std::vector<T> &arr) {
  MergeSort::sort(arr);
}

template <typename T> void quickSort(std::vector<T> &arr) {
  QuickSort::sort(arr);
}

template <typename T> void selectionSort(std::vector<T> &arr) {
  SelectionSort::sort(arr);
}
} // namespace Sorting

#endif
