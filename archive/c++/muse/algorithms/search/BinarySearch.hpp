#ifndef MUSE_ALGORITHMS_SEARCH_BINARY_SEARCH_HPP
#define MUSE_ALGORITHMS_SEARCH_BINARY_SEARCH_HPP

#include <cstddef>
#include <vector>

namespace BinarySearch {

template <typename T> std::size_t find(const std::vector<T> &arr, T key) {
  std::size_t lo = 0;
  std::size_t hi = arr.size();

  std::size_t mid;

  while (lo < hi) {
    mid = lo + ((hi - lo) >> 1);
    if (key < arr[mid]) {
      hi = mid;
    } else if (key > arr[mid]) {
      lo = mid + 1;
    } else {
      return mid;
    }
  }

  return arr.size();
}
} // namespace BinarySearch

#endif
