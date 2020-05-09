#ifndef FUN_VAC_ALGORITHMS_SORTING_MERGE_SORT_HPP
#define FUN_VAC_ALGORITHMS_SORTING_MERGE_SORT_HPP 1

#include <cstddef>
#include <vector>

namespace MergeSort {

    template <typename T>
    void merge(std::vector<T> &arr, std::size_t lo, std::size_t mid, std::size_t hi) {
        if (lo == mid) {
            return;
        }

        merge(arr, lo, (lo + mid) >> 1, mid);
        merge(arr, mid, (mid + hi) >> 1, hi);

        std::size_t m = lo;
        std::size_t n = mid;

        T *sorted = new T[hi - lo];

        for (std::size_t i = 0, length = hi - lo; i < length; i++) {
            if (m != mid && (n == hi || arr[m] < arr[n])) {
                sorted[i] = arr[m];
                m++;
            } else {
                sorted[i] = arr[n];
                n++;
            }
        }

        std::size_t cursor = 0;
        for (std::size_t i = lo; i < hi; i++) {
            arr[i] = sorted[cursor];
            cursor++;
        }

        delete[] sorted;
    }

    template <typename T>
    void sort(std::vector<T> &arr) {
        merge(arr, 0, arr.size() >> 1, arr.size());
    }
}

#endif
