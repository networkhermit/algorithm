#ifndef VAC_FUN_ALGORITHMS_META_SEARCH_HPP
#define VAC_FUN_ALGORITHMS_META_SEARCH_HPP 1

#include <vac.fun/algorithms/search/BinarySearch.hpp>
#include <vac.fun/algorithms/search/LinearSearch.hpp>

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
