#ifndef FUN_VAC_ALGORITHMS_SORTING_MERGE_SORT_H
#define FUN_VAC_ALGORITHMS_SORTING_MERGE_SORT_H 1

#include <stdlib.h>

#ifndef MERGE_SORT_TYPE
typedef int merge_sort_type;
#else
typedef MERGE_SORT_TYPE bubble_sort_type;
#endif

void MergeSort_merge(merge_sort_type *arr, size_t lo, size_t mid, size_t hi) {
    if (lo == mid) {
        return;
    }

    MergeSort_merge(arr, lo, (lo + mid) >> 1, mid);
    MergeSort_merge(arr, mid, (mid + hi) >> 1, hi);

    size_t m = lo;
    size_t n = mid;

    merge_sort_type *sorted = (merge_sort_type *) malloc(sizeof(merge_sort_type) * (hi - lo));

    for (size_t i = 0, length = hi - lo; i < length; i++) {
        if (m != mid && (n == hi || arr[m] < arr[n])) {
            sorted[i] = arr[m];
            m += 1;
        } else {
            sorted[i] = arr[n];
            n += 1;
        }
    }

    size_t cursor = 0;
    for (size_t i = lo; i < hi; i++) {
        arr[i] = sorted[cursor];
        cursor += 1;
    }

    free(sorted);
}

void MergeSort_sort(merge_sort_type *arr, size_t length) {
    MergeSort_merge(arr, 0, length >> 1, length);
}

#endif
