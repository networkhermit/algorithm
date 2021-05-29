#ifndef VAC_FUN_ALGORITHMS_SEARCH_BINARY_SEARCH_H
#define VAC_FUN_ALGORITHMS_SEARCH_BINARY_SEARCH_H 1

#include <stddef.h>

#ifndef BINARY_SEARCH_TYPE
typedef int binary_search_type;
#else
typedef BINARY_SEARCH_TYPE binary_search_type;
#endif

size_t BinarySearch_find(const binary_search_type *arr, size_t length, binary_search_type key) {
    size_t lo = 0;
    size_t hi = length;

    size_t mid;

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

    return length;
}

#endif
