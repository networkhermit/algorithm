#ifndef MUSE_ALGORITHMS_SORTING_QUICK_SORT_H
#define MUSE_ALGORITHMS_SORTING_QUICK_SORT_H

#include <stddef.h>

#ifndef QUICK_SORT_TYPE
typedef int quick_sort_type;
#else
typedef QUICK_SORT_TYPE quick_sort_type;
#endif

void QuickSort_partition(quick_sort_type *arr, size_t lo, size_t hi) {
    if (lo == hi) {
        return;
    }

    quick_sort_type pivot = arr[lo];

    size_t left = lo;
    size_t right = hi - 1;

    while (left != right) {
        for (size_t i = right; i > left; i--) {
            if (arr[i] < pivot) {
                arr[left] = arr[i];
                arr[i] = pivot;
                break;
            }
            right--;
        }
        for (size_t i = left; i < right; i++) {
            if (arr[i] > pivot) {
                arr[right] = arr[i];
                arr[i] = pivot;
                break;
            }
            left++;
        }
    }

    QuickSort_partition(arr, lo, left);
    QuickSort_partition(arr, left + 1, hi);
}

void QuickSort_sort(quick_sort_type *arr, size_t length) {
    QuickSort_partition(arr, 0, length);
}

#endif
