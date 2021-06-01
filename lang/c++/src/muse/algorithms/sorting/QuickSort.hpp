#ifndef MUSE_ALGORITHMS_SORTING_QUICK_SORT_HPP
#define MUSE_ALGORITHMS_SORTING_QUICK_SORT_HPP 1

#include <cstddef>
#include <vector>

namespace QuickSort {

    template <typename T>
    void partition(std::vector<T> &arr, std::size_t lo, std::size_t hi) {
        if (lo == hi) {
            return;
        }

        T pivot = arr[lo];

        std::size_t left = lo;
        std::size_t right = hi - 1;

        while (left != right) {
            for (std::size_t i = right; i > left; i--) {
                if (arr[i] < pivot) {
                    arr[left] = arr[i];
                    arr[i] = pivot;
                    break;
                }
                right--;
            }
            for (std::size_t i = left; i < right; i++) {
                if (arr[i] > pivot) {
                    arr[right] = arr[i];
                    arr[i] = pivot;
                    break;
                }
                left++;
            }
        }

        partition(arr, lo, left);
        partition(arr, left + 1, hi);
    }

    template <typename T>
    void sort(std::vector<T> &arr) {
        partition(arr, 0, arr.size());
    }
}

#endif
