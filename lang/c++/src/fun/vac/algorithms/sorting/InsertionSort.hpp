#ifndef FUN_VAC_ALGORITHMS_SORTING_INSERTION_SORT_HPP
#define FUN_VAC_ALGORITHMS_SORTING_INSERTION_SORT_HPP 1

#include <cstddef>
#include <vector>

namespace InsertionSort {

    template <typename T>
    void sort(std::vector<T> &arr) {
        T target;

        std::size_t cursor;

        for (std::size_t i = 1, length = arr.size(); i < length; i++) {
            target = arr[i];
            for (cursor = i; cursor > 0; cursor--) {
                if (arr[cursor - 1] <= target) {
                    break;
                }
                arr[cursor] = arr[cursor - 1];
            }
            arr[cursor] = target;
        }
    }
}

#endif
