#ifndef FUN_VAC_ALGORITHMS_SEARCH_LINEAR_SEARCH_H
#define FUN_VAC_ALGORITHMS_SEARCH_LINEAR_SEARCH_H 1

#include <stddef.h>

#ifndef LINEAR_SEARCH_TYPE
typedef int linear_search_type;
#else
typedef LINEAR_SEARCH_TYPE linear_search_type;
#endif

size_t LinearSearch_find(const linear_search_type *arr, size_t length, linear_search_type key) {
    for (size_t i = 0; i < length; i++) {
        if (arr[i] == key) {
            return i;
        }
    }

    return length;
}

#endif
