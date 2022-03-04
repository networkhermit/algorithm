#ifndef MUSE_ALGORITHMS_SORTING_BUBBLE_SORT_H
#define MUSE_ALGORITHMS_SORTING_BUBBLE_SORT_H

#include <stddef.h>

#ifndef BUBBLE_SORT_TYPE
typedef int bubble_sort_type;
#else
typedef BUBBLE_SORT_TYPE bubble_sort_type;
#endif

void BubbleSort_sort(bubble_sort_type *arr, size_t length) {
    bubble_sort_type temp;

    size_t margin;
    size_t unsorted = length;

    while (unsorted > 1) {
        margin = 0;
        for (size_t i = 1; i < unsorted; i++) {
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

#endif
