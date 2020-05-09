#ifndef FUN_VAC_ALGORITHMS_META_SEARCH_HPP
#define FUN_VAC_ALGORITHMS_META_SEARCH_HPP 1

#include <fun/vac/algorithms/search/BinarySearch.hpp>
#include <fun/vac/algorithms/search/LinearSearch.hpp>

namespace Search {

    template <typename T>
    std::size_t binarySearch(const std::vector<T> &arr, T key) {
        return BinarySearch::find(arr, key);
    }

    template <typename T>
    std::size_t linearSearch(const std::vector<T> &arr, T key) {
        return LinearSearch::find(arr, key);
    }
}

#endif
