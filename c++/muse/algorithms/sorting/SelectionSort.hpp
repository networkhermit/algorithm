#ifndef MUSE_ALGORITHMS_SORTING_SELECTION_SORT_HPP
#define MUSE_ALGORITHMS_SORTING_SELECTION_SORT_HPP

#include <cstddef>
#include <vector>

namespace SelectionSort {

template <typename T> void sort(std::vector<T> &arr) {
  if (arr.size() == 0) {
    return;
  }

  T temp;

  std::size_t iMin;

  for (std::size_t i = 0, bound = arr.size() - 1; i < bound; i++) {
    iMin = i;
    for (std::size_t j = i + 1; j <= bound; j++) {
      if (arr[j] < arr[iMin]) {
        iMin = j;
      }
    }
    if (iMin != i) {
      temp = arr[i];
      arr[i] = arr[iMin];
      arr[iMin] = temp;
    }
  }
}
} // namespace SelectionSort

#endif
