#ifndef MUSE_ALGORITHMS_SORTING_BUBBLE_SORT_HPP
#define MUSE_ALGORITHMS_SORTING_BUBBLE_SORT_HPP

#include <cstddef>
#include <vector>

namespace BubbleSort {

template <typename T> void sort(std::vector<T> &arr) {
  T temp;

  std::size_t margin;
  std::size_t unsorted = arr.size();

  while (unsorted > 1) {
    margin = 0;
    for (std::size_t i = 1; i < unsorted; i++) {
      if (arr[i - 1] > arr[i]) {
        temp = arr[i - 1];
        arr[i - 1] = arr[i];
        arr[i] = temp;
        margin = i;
      }
    }
    unsorted = margin;
  }
}
} // namespace BubbleSort

#endif
