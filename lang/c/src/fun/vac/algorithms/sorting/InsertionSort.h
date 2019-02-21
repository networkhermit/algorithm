#ifndef FUN_VAC_ALGORITHMS_SORTING_INSERTION_SORT_H
#define FUN_VAC_ALGORITHMS_SORTING_INSERTION_SORT_H 1

#include <stddef.h>

#ifndef INSERTION_SORT_TYPE
typedef int insertion_sort_type;
#else
typedef INSERTION_SORT_TYPE bubble_sort_type;
#endif

void InsertionSort_sort(insertion_sort_type *arr, size_t length) {
    insertion_sort_type target;

    size_t cursor;

    for (size_t i = 1; i < length; i++) {
        target = arr[i];
        for (cursor = i; cursor > 0; cursor--) {
            if (arr[cursor - 1] > target) {
                arr[cursor] = arr[cursor - 1];
            } else {
                break;
            }
        }
        arr[cursor] = target;
    }
}

#endif
