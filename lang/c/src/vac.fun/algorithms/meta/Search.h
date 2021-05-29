#ifndef VAC_FUN_ALGORITHMS_META_SEARCH_H
#define VAC_FUN_ALGORITHMS_META_SEARCH_H 1

#ifndef SEARCH_TYPE
typedef int search_type;
#else
typedef SEARCH_TYPE search_type;
#define BINARY_SEARCH_TYPE SEARCH_TYPE
#define LINEAR_SEARCH_TYPE SEARCH_TYPE
#endif

#include <vac.fun/algorithms/search/BinarySearch.h>
#include <vac.fun/algorithms/search/LinearSearch.h>

size_t Search_binarySearch(const search_type *arr, size_t length, search_type key) {
    return BinarySearch_find(arr, length, key);
}

size_t Search_linearSearch(const search_type *arr, size_t length, search_type key) {
    return LinearSearch_find(arr, length, key);
}

#endif
